/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ClubAthlete struct {
	// Resource state, indicates level of detail. Possible values: 1 -> \"meta\", 2 -> \"summary\", 3 -> \"detail\"
	ResourceState int32 `json:"resource_state,omitempty"`
	// The athlete's first name.
	Firstname string `json:"firstname,omitempty"`
	// The athlete's last initial.
	Lastname string `json:"lastname,omitempty"`
	// The athlete's member status.
	Member string `json:"member,omitempty"`
	// Whether the athlete is a club admin.
	Admin bool `json:"admin,omitempty"`
	// Whether the athlete is club owner.
	Owner bool `json:"owner,omitempty"`
}