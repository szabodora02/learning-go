package client

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"

    "kvstore/pkg/api"  // This is needed to use our own key-value store API.
)
// This is needed to use our own key-value store API.
type client struct {
	url string
}

func NewClient(addr string) Client {
    return &client{url: "http://" + addr}
}

// Get returns the value and version stored for the given key, or an error if something goes wrong.
func (c *client) Get(key string) (api.VersionedValue, error) {
   //this will append the query parameter formatted to the url
   uri, _ := url.Parse(c.url + "/api/get")
   q := uri.Query()
   q.Set("id", key)
   uri.RawQuery = q.Encode()
   // HTTP GET request
    resp, err := http.Get(uri.String())
    if err != nil {
        return api.VersionedValue{}, fmt.Errorf("failed to GET: %w", err)
    }
    defer resp.Body.Close()

    // handle server errors
    if resp.StatusCode != http.StatusOK {
        return api.VersionedValue{}, fmt.Errorf("server returned %s", resp.Status)
    }

    // decode response JSON into VersionedValue struct
    var vv api.VersionedValue
    if err := json.NewDecoder(resp.Body).Decode(&vv); err != nil {
        return api.VersionedValue{}, fmt.Errorf("failed to decode response: %w", err)
    }

    return vv, nil
}

// Put tries to insert the given key-value pair with the specified version into the store.
func (c *client) Put(vkv api.VersionedKeyValue) error {

    body, err := json.Marshal(vkv)
    if err != nil {
        return fmt.Errorf("failed to marshal key-value: %w", err)
    }

    resp, err := http.Post(c.url+"/api/put", "application/json", bytes.NewBuffer(body))
    if err != nil {
        return fmt.Errorf("failed to POST: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("server returned %s", resp.Status)
    }

    return nil
}

// List returns all values stored in the database.
func (c *client) List() ([]api.VersionedKeyValue, error) {
    resp, err := http.Get(c.url + "/api/list")
    if err != nil {
        return []api.VersionedKeyValue{}, fmt.Errorf("failed to GET: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return []api.VersionedKeyValue{}, fmt.Errorf("server returned %s", resp.Status)
    }

    var kvList []api.VersionedKeyValue
    if err := json.NewDecoder(resp.Body).Decode(&kvList); err != nil {
        return []api.VersionedKeyValue{}, fmt.Errorf("failed to decode response: %w", err)
    }

    return kvList, nil
}
// Reset removes all key-value pairs.
func (c *client) Reset() error {
    resp, err := http.Get(c.url + "/api/reset")
    if err != nil {
        return fmt.Errorf("failed to GET reset: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("server returned %s", resp.Status)
    }

    return nil
}
