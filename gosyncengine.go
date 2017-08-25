package gosyncengine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

var (
	threadsType     = reflect.TypeOf(new(Threads))
	threadType      = reflect.TypeOf(new(Thread))
	deltaCursorType = reflect.TypeOf(new(DeltaCursor))
)

// SyncEngineAPI provides access to the sync engine API
type SyncEngineAPI struct {
	BaseURL string
}

// New creates a new SyncEngine API object
func New(baseURL string) *SyncEngineAPI {
	return &SyncEngineAPI{
		BaseURL: baseURL,
	}
}

func (api *SyncEngineAPI) getURL(path string) string {
	return fmt.Sprintf("%s%s", api.BaseURL, path)
}

func (api *SyncEngineAPI) executeRequest(method string, userID string, path string, requestBody []byte) (*http.Response, error) {
	var requestBuffer io.Reader
	if requestBody != nil {
		requestBuffer = bytes.NewBuffer(requestBody)
	}

	var url = api.getURL(path)
	req, _ := http.NewRequest(method, url, requestBuffer)
	req.SetBasicAuth(userID, "")

	client := &http.Client{}

	var resp *http.Response
	var err error

	if resp, err = client.Do(req); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAccounts returns all the accounts defined on the server
func (api *SyncEngineAPI) GetAccounts() (Accounts, error) {
	var req *http.Request
	var resp *http.Response
	var err error
	req, _ = http.NewRequest(http.MethodGet, api.getURL("/accounts"), nil)

	client := &http.Client{}
	if resp, err = client.Do(req); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Reading request body failed. Status=%d  Reason=%s", resp.StatusCode, string(body))
	}

	var result Accounts
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Response deserialization failed. Reason: %s", err)
	}

	return result, nil
}

// GetThreads returns all the threads of the specified account ID
func (api *SyncEngineAPI) GetThreads(accountID string) (Threads, error) {
	var resp *http.Response
	var err error

	if resp, err = api.executeRequest(http.MethodGet, accountID, "/threads", nil); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Reading request body failed. Status=%d  Reason=%s", resp.StatusCode, string(body))
	}

	var result Threads
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Response deserialization failed. Reason: %s", err)
	}

	return result, nil
}

// GetThreadByID returns a thread by its ID
func (api *SyncEngineAPI) GetThreadByID(accountID string, threadID string) (*Thread, error) {
	var resp *http.Response
	var err error

	if resp, err = api.executeRequest(http.MethodGet, accountID, fmt.Sprintf("/threads/%s", threadID), nil); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Reading request body failed. Status=%d  Reason=%s", resp.StatusCode, string(body))
	}

	var result = &Thread{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Response deserialization failed. Reason: %s", err)
	}

	return result, nil
}

// GetMessageByID returns a single message by its ID
func (api *SyncEngineAPI) GetMessageByID(accountID string, messageID string) (*Message, error) {
	var resp *http.Response
	var err error

	if resp, err = api.executeRequest(http.MethodGet, accountID, fmt.Sprintf("/messages/%s", messageID), nil); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Reading request body failed. Status=%d  Reason=%s", resp.StatusCode, string(body))
	}

	var result = &Message{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Response deserialization failed. Reason: %s", err)
	}

	return result, nil
}

// GetThreadMessages returns all of the messages associated with the specified thread ID
func (api *SyncEngineAPI) GetThreadMessages(accountID string, threadID string) (Messages, error) {
	var resp *http.Response
	var err error

	if resp, err = api.executeRequest(http.MethodGet, accountID, fmt.Sprintf("/messages/?thread_id=%s", threadID), nil); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Reading request body failed. Status=%d  Reason=%s", resp.StatusCode, string(body))
	}

	var result = Messages{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Response deserialization failed. Reason: %s", err)
	}

	return result, nil
}

// GetDeltaLatestCursor returns the latest cursor available
func (api *SyncEngineAPI) GetDeltaLatestCursor(accountID string) (*DeltaCursor, error) {
	var resp *http.Response
	var err error

	if resp, err = api.executeRequest(http.MethodPost, accountID, "/delta/latest_cursor", nil); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Reading request body failed. Status=%d  Reason=%s", resp.StatusCode, string(body))
	}

	var result = &DeltaCursor{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Response deserialization failed. Reason: %s", err)
	}

	return result, nil
}

// GetDeltaMessages will return only messages from the given cursor
func (api *SyncEngineAPI) GetDeltaMessages(accountID string, cursor string) (*DeltaMessages, error) {
	var resp *http.Response
	var err error

	if resp, err = api.executeRequest(http.MethodGet, accountID, fmt.Sprintf("/delta?cursor=%s&view=expanded&include_types=message", cursor), nil); err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Reading request body failed. Status=%d  Reason=%s", resp.StatusCode, string(body))
	}

	var result = &DeltaMessages{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Response deserialization failed. Reason: %s", err)
	}

	return result, nil
}
