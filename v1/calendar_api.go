/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev7
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package goesiv1

import (
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type CalendarApiService service


/* CalendarApiService List calendar event summaries
 Get 50 event summaries from the calendar. If no event ID is given, the resource will return the next 50 chronological event summaries from now. If an event ID is specified, it will return the next 50 chronological event summaries from after that event.   ---  Alternate route: &#x60;/legacy/characters/{character_id}/calendar/&#x60;  Alternate route: &#x60;/latest/characters/{character_id}/calendar/&#x60;  Alternate route: &#x60;/dev/characters/{character_id}/calendar/&#x60;   ---  This route is cached for up to 5 seconds

 * @param ctx context.Context Authentication Context 
 @param characterId The character to retrieve events from
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "datasource" (string) The server name you would like data from
     @param "fromEvent" (int32) The event ID to retrieve events from
     @param "token" (string) Access token to use, if preferred over a header
     @param "userAgent" (string) Client identifier, takes precedence over headers
     @param "xUserAgent" (string) Client identifier, takes precedence over User-Agent
 @return []GetCharactersCharacterIdCalendar200Ok*/
func (a *CalendarApiService) GetCharactersCharacterIdCalendar(ctx context.Context, characterId int32, localVarOptionals map[string]interface{}) ([]GetCharactersCharacterIdCalendar200Ok,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []GetCharactersCharacterIdCalendar200Ok
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/characters/{character_id}/calendar/"
	localVarPath = strings.Replace(localVarPath, "{"+"character_id"+"}", fmt.Sprintf("%v", characterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["datasource"], "string", "datasource"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fromEvent"], "int32", "fromEvent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["token"], "string", "token"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["userAgent"], "string", "userAgent"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xUserAgent"], "string", "xUserAgent"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["datasource"].(string); localVarOk {
		localVarQueryParams.Add("datasource", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fromEvent"].(int32); localVarOk {
		localVarQueryParams.Add("from_event", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["token"].(string); localVarOk {
		localVarQueryParams.Add("token", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["userAgent"].(string); localVarOk {
		localVarQueryParams.Add("user_agent", parameterToString(localVarTempParam, ""))
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xUserAgent"].(string); localVarOk {
		localVarHeaderParams["X-User-Agent"] = parameterToString(localVarTempParam, "")
	}

	 r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	 if err != nil {
		  return successPayload, nil, err
	 }

	 localVarHttpResponse, err := a.client.callAPI(r)
	 if err != nil || localVarHttpResponse == nil {
		  return successPayload, localVarHttpResponse, err
	 }
	 defer localVarHttpResponse.Body.Close()
	 if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	 }
	
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
	 	return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

