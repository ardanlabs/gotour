// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample test to show how to write a basic unit test.
package example1

import (
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestDownload validates the http Get function can download content.
func TestDownload(t *testing.T) {
	url := "https://www.ardanlabs.com/blog/index.xml"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen checking %q for status code %d", testID, url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to make the Get call : %v", failed, testID, err)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to make the Get call.", succeed, testID)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t%s\tTest %d:\tShould receive a %d status code.", succeed, testID, statusCode)
			} else {
				t.Errorf("\t%s\tTest %d:\tShould receive a %d status code : %d", failed, testID, statusCode, resp.StatusCode)
			}
		}
	}
}
