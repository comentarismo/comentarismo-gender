package server_test

import (
	"comentarismo-gender/gender"
	"comentarismo-gender/server"
	"encoding/json"
	"github.com/drewolson/testflight"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestMaleGenderHandler(t *testing.T) {

	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {

		Convey("Should Learn female names in english and report gender female for a female name", t, func() {
			targetWord := "matt"
			response := r.Post("/report?lang=en&gender=M", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "matthew"
			response = r.Post("/report?lang=en&gender=M", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "hank"
			response = r.Post("/report?lang=en&gender=M", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "mark"
			response = r.Post("/report?lang=en&gender=M", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "edward"
			response = r.Post("/report?lang=en&gender=M", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "henry"
			response = r.Post("/report?lang=en&gender=M", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "charlie"
			response = r.Post("/report?lang=en&gender=M", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "ben"
			response = r.Post("/report?lang=en&gender=M", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			//now try with a name
			textTarget := "charlie"
			response = r.Post("/gender?lang=en", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","spam":false}

			genderReport := gender.GenderReport{}
			err := json.Unmarshal(response.RawBody, &genderReport)
			So(err, ShouldBeNil)

			So(genderReport.Error, ShouldBeBlank)
			So(genderReport.Code, ShouldEqual, 200)
			So(genderReport.Gender, ShouldEqual, "male")

			//now try with another male name
			textTarget = "henry"
			response = r.Post("/gender?lang=en", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","gender":true}

			genderReport = gender.GenderReport{}
			err = json.Unmarshal(response.RawBody, &genderReport)
			So(err, ShouldBeNil)

			So(genderReport.Error, ShouldBeBlank)
			So(genderReport.Code, ShouldEqual, 200)
			So(genderReport.Gender, ShouldEqual, "male")

			//now revoke the name
			textTarget = "henry"
			response = r.Post("/revoke?lang=en&gender=F", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)
			log.Println(response.Body) //{"code":200,"error":"","gender":true}

			//now it should not be that gender anymore
			textTarget = "henry"
			response = r.Post("/gender?lang=en", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body)

			genderReport = gender.GenderReport{}
			err = json.Unmarshal(response.RawBody, &genderReport)
			So(err, ShouldBeNil)

			So(genderReport.Error, ShouldBeBlank)
			So(genderReport.Code, ShouldEqual, 200)
			So(genderReport.Gender, ShouldEqual, "male")
		})

	})
}

func init() {
	gender.Flush()
}
