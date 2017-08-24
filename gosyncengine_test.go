package gosyncengine

import (
	"os"
	"strings"
	"testing"

	"github.com/maxcnunes/httpfake"
)

var (
	envCache = map[string]string{}
)

func getEnvValue(key string) string {
	if len(envCache) == 0 {
		envData := os.Environ()
		for _, v := range envData {
			parts := strings.Split(v, "=")
			envCache[parts[0]] = parts[1]
		}
	}

	if v, ok := envCache[key]; ok {
		return v
	}

	return ""
}

func TestGetAccountBadURL(t *testing.T) {
	client := New("aaaaaa")

	if _, err := client.GetAccounts(); err == nil {
		t.Error(err)
	}
}

func TestGetAccounts(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/accounts").
		Reply(200).
		BodyString(`[
    {
        "account_id": "11111",
        "email_address": "a@b.com",
        "id": "xxxx",
        "name": "",
        "object": "account",
        "organization_unit": "folder",
        "provider": "custom",
        "sync_state": "running"
    },
    {
        "account_id": "22222",
        "email_address": "a@c.com",
        "id": "yyyy",
        "name": "",
        "object": "account",
        "organization_unit": "folder",
        "provider": "custom",
        "sync_state": "running"
    }]`)

	client := New(fakeService.ResolveURL(""))
	var accounts []Account
	var err error

	if accounts, err = client.GetAccounts(); err != nil {
		t.Error(err)
	} else {
		t.Logf("GetAccounts Result: %v", accounts)
	}
}

func TestGetAccountsNot200(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/accounts").
		Reply(404)

	client := New(fakeService.ResolveURL(""))
	if _, err := client.GetAccounts(); err == nil {
		t.Error("Should have gotten an error here")
	}
}

func TestGetAccountsBadUnmarshaling(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/accounts").
		Reply(200).
		BodyString(`a]{`)

	client := New(fakeService.ResolveURL(""))
	if _, err := client.GetAccounts(); err == nil {
		t.Error("Should have gotten a bad JSON response and failed")
	}
}

func TestGetThreads(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/threads").
		Reply(200).
		BodyString(`[
    {
        "account_id": "5qro12wr9mojq8y9y8f6cdcp",
        "draft_ids": [],
        "first_message_timestamp": 1500437315,
        "folders": [
            {
                "display_name": "INBOX",
                "id": "a9n0r8jfqc8v6vih7tfeykve5",
                "name": "inbox"
            }
        ],
        "has_attachments": false,
        "id": "6w6csoo143g2r0a8g0gppyuhw",
        "last_message_received_timestamp": 1500437315,
        "last_message_sent_timestamp": null,
        "last_message_timestamp": 1500437315,
        "message_ids": [
            "9oqpu6tm1kzg09yr080i6498m"
        ],
        "object": "thread",
        "participants": [
            {
                "email": "mail-noreply@google.com",
                "name": "Gmail Team"
            },
            {
                "email": "cindy@morefblike.com",
                "name": "Cindy White"
            }
        ],
        "snippet": "Hi CindyWelcome to your Gmail inbox Save everything With tons of storage space, you\u2019ll never need to delete an email. Just keep everything and easily find it later. Find emails fast With the ",
        "starred": false,
        "subject": "Tips for using your new inbox",
        "unread": true,
        "version": 1
    },
    {
        "account_id": "5qro12wr9mojq8y9y8f6cdcp",
        "draft_ids": [],
        "first_message_timestamp": 1500437314,
        "folders": [
            {
                "display_name": "[Gmail]/All Mail",
                "id": "917ucyqu8q9rmabkax89v3wyi",
                "name": null
            }
        ],
        "has_attachments": false,
        "id": "a28swc983tzc5ledvc3550glo",
        "last_message_received_timestamp": 1500437314,
        "last_message_sent_timestamp": null,
        "last_message_timestamp": 1500437314,
        "message_ids": [
            "a1sf26fz65atp7anh7qhxfh9m"
        ],
        "object": "thread",
        "participants": [
            {
                "email": "mail-noreply@google.com",
                "name": "Gmail Team"
            },
            {
                "email": "cindy@morefblike.com",
                "name": "Cindy White"
            }
        ],
        "snippet": "Hi Cindy Get the official Gmail app The best features of Gmail are only available on your phone and tablet with the official Gmail app. Download the app or go to gmail.com on your computer or",
        "starred": false,
        "subject": "The best of Gmail, wherever you are",
        "unread": true,
        "version": 1
    }]`)

	client := New(fakeService.ResolveURL(""))
	var threads Threads
	var err error

	if threads, err = client.GetThreads("aaa"); err != nil {
		t.Error(err)
	} else {
		t.Logf("GetThreads Result: %v", threads)
	}
}

func TestGetThreadsNot200(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/threads").
		Reply(404)

	client := New(fakeService.ResolveURL(""))

	if _, err := client.GetThreads("aaa"); err == nil {
		t.Error(err)
	}
}

func TestGetThreadsBadUnmarshaling(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/threads").
		Reply(200).
		BodyString(`a]{`)

	client := New(fakeService.ResolveURL(""))
	if _, err := client.GetThreads("aaa"); err == nil {
		t.Error("Should have gotten a bad JSON response and failed")
	}
}

func TestGetThreadByID(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/threads/aaa").
		Reply(200).
		BodyString(`{
    "account_id": "5qro12wr9mojq8y9y8f6cdcp",
    "draft_ids": [],
    "first_message_timestamp": 1500437314,
    "folders": [
        {
            "display_name": "[Gmail]/All Mail",
            "id": "917ucyqu8q9rmabkax89v3wyi",
            "name": null
        }
    ],
    "has_attachments": false,
    "id": "a28swc983tzc5ledvc3550glo",
    "last_message_received_timestamp": 1500437314,
    "last_message_sent_timestamp": null,
    "last_message_timestamp": 1500437314,
    "message_ids": [
        "a1sf26fz65atp7anh7qhxfh9m"
    ],
    "object": "thread",
    "participants": [
        {
            "email": "mail-noreply@google.com",
            "name": "Gmail Team"
        },
        {
            "email": "cindy@morefblike.com",
            "name": "Cindy White"
        }
    ],
    "snippet": "Hi Cindy Get the official Gmail app The best features of Gmail are only available on your phone and tablet with the official Gmail app. Download the app or go to gmail.com on your computer or",
    "starred": false,
    "subject": "The best of Gmail, wherever you are",
    "unread": true,
    "version": 1
	}`)

	client := New(fakeService.ResolveURL(""))
	var thread *Thread
	var err error

	if thread, err = client.GetThreadByID("aaa", "aaa"); err != nil {
		t.Error(err)
	} else {
		t.Logf("GetThreadByID Result: %v", thread)
	}
}

func TestGetThreadByIDNot200(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/thread/aaa").
		Reply(404)

	client := New(fakeService.ResolveURL(""))

	if _, err := client.GetThreadByID("aaa", "aaa"); err == nil {
		t.Error(err)
	}
}

func TestGetThreadByIDBadUnmarshaling(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/threads/aaa").
		Reply(200).
		BodyString(`a]{`)

	client := New(fakeService.ResolveURL(""))
	if _, err := client.GetThreadByID("aaa", "aaa"); err == nil {
		t.Error("Should have gotten a bad JSON response and failed")
	}
}

func TestGetMessageByID(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/messages/aaa").
		Reply(200).
		BodyString(`{
        "account_id": "zzz",
        "bcc": [],
        "body": "<!DOCTYPE html>\n<html><head><body>Hello World</body></html>",
        "cc": [],
        "date": 1500437314,
        "events": [],
        "files": [],
        "folder": {
            "display_name": "All Mail"
				},
        "from": [
            {
                "email": "mail-noreply@somewhere.com",
                "name": "Team"
            }
        ],
        "id": "aaa",
        "object": "message",
        "reply_to": [],
        "snippet": "das djaskl djsakld jaskldj sadklasjdklasjdaklsdjaskl",
        "starred": false,
        "subject": "The best Email. Ever.",
        "thread_id": "bbb",
        "to": [
            {
                "email": "a@b.com",
                "name": "a b"
            }
        ],
        "unread": true
    }`)

	client := New(fakeService.ResolveURL(""))
	var message *Message
	var err error

	if message, err = client.GetMessageByID("aaa", "aaa"); err != nil {
		t.Error(err)
	} else {
		t.Logf("GetMessageByID Result: %v", message)
	}
}

func TestGetMessageByIDNot200(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/messages/aaa").
		Reply(404)

	client := New(fakeService.ResolveURL(""))

	if _, err := client.GetMessageByID("aaa", "aaa"); err == nil {
		t.Error(err)
	}
}

func TestGetMessageByIDBadUnmarshaling(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/messages/aaa").
		Reply(200).
		BodyString(`a]{`)

	client := New(fakeService.ResolveURL(""))
	if _, err := client.GetMessageByID("aaa", "aaa"); err == nil {
		t.Error("Should have gotten a bad JSON response and failed")
	}
}

func TestGetThreadMessages(t *testing.T) {
	baseURL := getEnvValue("SYNCENGINE_URL")
	client := New(baseURL)

	var messages Messages
	var err error
	if messages, err = client.GetThreadMessages("5qro12wr9mojq8y9y8f6cdcp", "a28swc983tzc5ledvc3550glo"); err != nil {
		t.Error(err)
	} else {
		t.Logf("GetThreadMessages: %v", messages)
		t.Logf("GetThreadMessages Count: %d", len(messages))
	}
}

func TestGetThreadMessagesNot200(t *testing.T) {
	fakeService := httpfake.New()
	defer fakeService.Server.Close()

	fakeService.NewHandler().
		Get("/messages/aaa").
		Reply(404)

	client := New(fakeService.ResolveURL(""))

	if _, err := client.GetThreadMessages("aaa", "aaa"); err == nil {
		t.Error(err)
	}
}
