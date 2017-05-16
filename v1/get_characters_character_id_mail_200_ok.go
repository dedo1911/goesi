/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.6.dev11
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
	"time"
)

/* 200 ok object */
type GetCharactersCharacterIdMail200Ok struct {
	From       int32                                   `json:"from,omitempty"`       /* From whom the mail was sent */
	IsRead     bool                                    `json:"is_read,omitempty"`    /* is_read boolean */
	Labels     []int64                                 `json:"labels,omitempty"`     /* labels array */
	MailId     int64                                   `json:"mail_id,omitempty"`    /* mail_id integer */
	Recipients []GetCharactersCharacterIdMailRecipient `json:"recipients,omitempty"` /* Recipients of the mail */
	Subject    string                                  `json:"subject,omitempty"`    /* Mail subject */
	Timestamp  time.Time                               `json:"timestamp,omitempty"`  /* When the mail was sent */
}
