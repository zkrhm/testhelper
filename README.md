# THE TEST HELPER

stop repeating yourself writing bootstrap codes and use this helper.

usage : 

1. create httptest struct calls the NewHttpTest constructor
2. execute DoRequestTest and get the response recorder object and err.
3. do matching to the responses.
   
```go

test := testhelper.NewHttpTest("POST", "/hello", "", app.Hello)
		rr, err := test.DoRequestTest()

		Expect(err).ShouldNot(HaveOccurred())
		Expect(rr.Code).Should(Equal(200))
		Expect(rr.Header().Get("Content-Type")).Should(Equal("application/json"))
        Expect(ioutil.ReadAll(rr.Body)).Should(MatchJSON(`{"code":200,"message":"Hello"}`))
        
```