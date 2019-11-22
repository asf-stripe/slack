package slack

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockConversion(t *testing.T) {
	data := []byte(`[{
		"type": "call",
		"call_id": "ABCDEFGHI",
		"block_id": "jHZ6",
		"api_decoration_available": false,
		"call": {
		    "v1": {
			"id": "ABCDEFGHI",
			"app_id": "ABCDEFGHI",
			"app_icon_urls": {
			    "image_32": "some-url.png",
			    "image_36": "some-url.png",
			    "image_48": "some-url.png",
			    "image_64": "some-url.png",
			    "image_72": "some-url.png",
			    "image_96": "some-url.png",
			    "image_128": "some-url.png",
			    "image_192": "some-url.png",
			    "image_512": "some-url.png",
			    "image_1024": "some-url.png",
			    "image_original": "some-url.png"
			},
			"date_start": 1566248229,
			"active_participants": [],
			"all_participants": [
			    {
				"slack_id": "ABCDEFGHI"
			    },
			    {
				"external_id": "12345678",
				"avatar_url": "",
				"display_name": "John Smith"
			    }
			],
			"display_id": "123-456-789",
			"join_url": "https:\/\/blah.zoom.us\/j\/123456789",
			"name": "Zoom meeting started by smith",
			"created_by": "ABCDEFGHI",
			"date_end": 1566249369,
			"channels": [
			    "ABCDEFGHI"
			],
			"is_dm_call": false,
			"was_rejected": false,
			"was_missed": false,
			"was_accepted": false,
			"has_ended": true
		    }
		}
	    }]`)
	var blocks Blocks
	err := json.Unmarshal(data, &blocks)
	assert.NoError(t, err)
	if assert.Equal(t, 1, len(blocks.BlockSet)) {
		block := blocks.BlockSet[0].(*UnknownBlock)
		assert.Equal(t, MessageBlockType("call"), block.BlockType())
	}

}
