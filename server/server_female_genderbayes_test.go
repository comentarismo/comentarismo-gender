package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewolson/testflight"
	"testing"
	"log"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"comentarismo-gender/server"
	"comentarismo-gender/gender"
)

func TestFemaleGenderHandler(t *testing.T) {

	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {

		Convey("Should Learn female names in english and report gender female for a female name", t, func() {
			targetWord := "rachael"
			response := r.Post("/report?lang=en&gender=F", testflight.FORM_ENCODED, "text=" + targetWord);
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "hannah"
			response = r.Post("/report?lang=en&gender=F", testflight.FORM_ENCODED, "text=" + targetWord);
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "norah"
			response = r.Post("/report?lang=en&gender=F", testflight.FORM_ENCODED, "text=" + targetWord);
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "claire"
			response = r.Post("/report?lang=en&gender=F", testflight.FORM_ENCODED, "text=" + targetWord);
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "marsha"
			response = r.Post("/report?lang=en&gender=F", testflight.FORM_ENCODED, "text=" + targetWord);
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "rachael"
			response = r.Post("/report?lang=en&gender=F", testflight.FORM_ENCODED, "text=" + targetWord);
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)


			//now try with a name
			textTarget := "rachael"
			response = r.Post("/gender?lang=en", testflight.FORM_ENCODED, "text="+textTarget );

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","spam":false}

			genderReport := gender.GenderReport{}
			err := json.Unmarshal(response.RawBody, &genderReport)
			So(err, ShouldBeNil)

			So(genderReport.Error, ShouldBeBlank)
			So(genderReport.Code, ShouldEqual, 200)
			So(genderReport.Gender, ShouldEqual, "female")

			//now try with a spammy comment
			textTarget = "marsha"
			response = r.Post("/gender?lang=en", testflight.FORM_ENCODED, "text="+textTarget );

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","gender":true}

			genderReport = gender.GenderReport{}
			err = json.Unmarshal(response.RawBody, &genderReport)
			So(err, ShouldBeNil)

			So(genderReport.Error, ShouldBeBlank)
			So(genderReport.Code, ShouldEqual, 200)
			So(genderReport.Gender, ShouldEqual, "female")


			//now revoke the gendermy comment
			textTarget = "marsha"
			response = r.Post("/revoke?lang=en&gender=F", testflight.FORM_ENCODED, "text="+textTarget );

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)
			log.Println(response.Body) //{"code":200,"error":"","gender":true}

			//now it should not be gender anymore
			textTarget = "marsha"
			response = r.Post("/gender?lang=en", testflight.FORM_ENCODED, "text="+textTarget );

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body)

			genderReport = gender.GenderReport{}
			err = json.Unmarshal(response.RawBody, &genderReport)
			So(err, ShouldBeNil)

			So(genderReport.Error, ShouldBeBlank)
			So(genderReport.Code, ShouldEqual, 200)
			So(genderReport.Gender, ShouldEqual, "female")
		})

	})
}

func init() {
	gender.Flush()
}
