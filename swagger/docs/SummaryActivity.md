# SummaryActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of the activity | [optional] 
**ExternalId** | Pointer to **string** | The identifier provided at upload time | [optional] 
**UploadId** | Pointer to **int64** | The identifier of the upload that resulted in this activity | [optional] 
**Athlete** | Pointer to [**MetaAthlete**](MetaAthlete.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the activity | [optional] 
**Distance** | Pointer to **float32** | The activity&#39;s distance, in meters | [optional] 
**MovingTime** | Pointer to **int32** | The activity&#39;s moving time, in seconds | [optional] 
**ElapsedTime** | Pointer to **int32** | The activity&#39;s elapsed time, in seconds | [optional] 
**TotalElevationGain** | Pointer to **float32** | The activity&#39;s total elevation gain. | [optional] 
**ElevHigh** | Pointer to **float32** | The activity&#39;s highest elevation, in meters | [optional] 
**ElevLow** | Pointer to **float32** | The activity&#39;s lowest elevation, in meters | [optional] 
**Type** | Pointer to [**ActivityType**](ActivityType.md) |  | [optional] 
**SportType** | Pointer to [**SportType**](SportType.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** | The time at which the activity was started. | [optional] 
**StartDateLocal** | Pointer to **time.Time** | The time at which the activity was started in the local timezone. | [optional] 
**Timezone** | Pointer to **string** | The timezone of the activity | [optional] 
**StartLatlng** | Pointer to **[]float32** | A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers. | [optional] 
**EndLatlng** | Pointer to **[]float32** | A pair of latitude/longitude coordinates, represented as an array of 2 floating point numbers. | [optional] 
**AchievementCount** | Pointer to **int32** | The number of achievements gained during this activity | [optional] 
**KudosCount** | Pointer to **int32** | The number of kudos given for this activity | [optional] 
**CommentCount** | Pointer to **int32** | The number of comments for this activity | [optional] 
**AthleteCount** | Pointer to **int32** | The number of athletes for taking part in a group activity | [optional] 
**PhotoCount** | Pointer to **int32** | The number of Instagram photos for this activity | [optional] 
**TotalPhotoCount** | Pointer to **int32** | The number of Instagram and Strava photos for this activity | [optional] 
**Map** | Pointer to [**PolylineMap**](PolylineMap.md) |  | [optional] 
**DeviceName** | Pointer to **string** | The name of the device used to record the activity | [optional] 
**Trainer** | Pointer to **bool** | Whether this activity was recorded on a training machine | [optional] 
**Commute** | Pointer to **bool** | Whether this activity is a commute | [optional] 
**Manual** | Pointer to **bool** | Whether this activity was created manually | [optional] 
**Private** | Pointer to **bool** | Whether this activity is private | [optional] 
**Flagged** | Pointer to **bool** | Whether this activity is flagged | [optional] 
**WorkoutType** | Pointer to **int32** | The activity&#39;s workout type | [optional] 
**UploadIdStr** | Pointer to **string** | The unique identifier of the upload in string format | [optional] 
**AverageSpeed** | Pointer to **float32** | The activity&#39;s average speed, in meters per second | [optional] 
**MaxSpeed** | Pointer to **float32** | The activity&#39;s max speed, in meters per second | [optional] 
**HasKudoed** | Pointer to **bool** | Whether the logged-in athlete has kudoed this activity | [optional] 
**HideFromHome** | Pointer to **bool** | Whether the activity is muted | [optional] 
**GearId** | Pointer to **string** | The id of the gear for the activity | [optional] 
**Kilojoules** | Pointer to **float32** | The total work done in kilojoules during this activity. Rides only | [optional] 
**AverageWatts** | Pointer to **float32** | Average power output in watts during this activity. Rides only | [optional] 
**DeviceWatts** | Pointer to **bool** | Whether the watts are from a power meter, false if estimated | [optional] 
**MaxWatts** | Pointer to **int32** | Rides with power meter data only | [optional] 
**WeightedAverageWatts** | Pointer to **int32** | Similar to Normalized Power. Rides with power meter data only | [optional] 

## Methods

### NewSummaryActivity

`func NewSummaryActivity() *SummaryActivity`

NewSummaryActivity instantiates a new SummaryActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryActivityWithDefaults

`func NewSummaryActivityWithDefaults() *SummaryActivity`

NewSummaryActivityWithDefaults instantiates a new SummaryActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SummaryActivity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SummaryActivity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SummaryActivity) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SummaryActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *SummaryActivity) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SummaryActivity) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SummaryActivity) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SummaryActivity) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUploadId

`func (o *SummaryActivity) GetUploadId() int64`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *SummaryActivity) GetUploadIdOk() (*int64, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *SummaryActivity) SetUploadId(v int64)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *SummaryActivity) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetAthlete

`func (o *SummaryActivity) GetAthlete() MetaAthlete`

GetAthlete returns the Athlete field if non-nil, zero value otherwise.

### GetAthleteOk

`func (o *SummaryActivity) GetAthleteOk() (*MetaAthlete, bool)`

GetAthleteOk returns a tuple with the Athlete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthlete

`func (o *SummaryActivity) SetAthlete(v MetaAthlete)`

SetAthlete sets Athlete field to given value.

### HasAthlete

`func (o *SummaryActivity) HasAthlete() bool`

HasAthlete returns a boolean if a field has been set.

### GetName

`func (o *SummaryActivity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SummaryActivity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SummaryActivity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SummaryActivity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistance

`func (o *SummaryActivity) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *SummaryActivity) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *SummaryActivity) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *SummaryActivity) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetMovingTime

`func (o *SummaryActivity) GetMovingTime() int32`

GetMovingTime returns the MovingTime field if non-nil, zero value otherwise.

### GetMovingTimeOk

`func (o *SummaryActivity) GetMovingTimeOk() (*int32, bool)`

GetMovingTimeOk returns a tuple with the MovingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovingTime

`func (o *SummaryActivity) SetMovingTime(v int32)`

SetMovingTime sets MovingTime field to given value.

### HasMovingTime

`func (o *SummaryActivity) HasMovingTime() bool`

HasMovingTime returns a boolean if a field has been set.

### GetElapsedTime

`func (o *SummaryActivity) GetElapsedTime() int32`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *SummaryActivity) GetElapsedTimeOk() (*int32, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *SummaryActivity) SetElapsedTime(v int32)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *SummaryActivity) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetTotalElevationGain

`func (o *SummaryActivity) GetTotalElevationGain() float32`

GetTotalElevationGain returns the TotalElevationGain field if non-nil, zero value otherwise.

### GetTotalElevationGainOk

`func (o *SummaryActivity) GetTotalElevationGainOk() (*float32, bool)`

GetTotalElevationGainOk returns a tuple with the TotalElevationGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElevationGain

`func (o *SummaryActivity) SetTotalElevationGain(v float32)`

SetTotalElevationGain sets TotalElevationGain field to given value.

### HasTotalElevationGain

`func (o *SummaryActivity) HasTotalElevationGain() bool`

HasTotalElevationGain returns a boolean if a field has been set.

### GetElevHigh

`func (o *SummaryActivity) GetElevHigh() float32`

GetElevHigh returns the ElevHigh field if non-nil, zero value otherwise.

### GetElevHighOk

`func (o *SummaryActivity) GetElevHighOk() (*float32, bool)`

GetElevHighOk returns a tuple with the ElevHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevHigh

`func (o *SummaryActivity) SetElevHigh(v float32)`

SetElevHigh sets ElevHigh field to given value.

### HasElevHigh

`func (o *SummaryActivity) HasElevHigh() bool`

HasElevHigh returns a boolean if a field has been set.

### GetElevLow

`func (o *SummaryActivity) GetElevLow() float32`

GetElevLow returns the ElevLow field if non-nil, zero value otherwise.

### GetElevLowOk

`func (o *SummaryActivity) GetElevLowOk() (*float32, bool)`

GetElevLowOk returns a tuple with the ElevLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevLow

`func (o *SummaryActivity) SetElevLow(v float32)`

SetElevLow sets ElevLow field to given value.

### HasElevLow

`func (o *SummaryActivity) HasElevLow() bool`

HasElevLow returns a boolean if a field has been set.

### GetType

`func (o *SummaryActivity) GetType() ActivityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SummaryActivity) GetTypeOk() (*ActivityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SummaryActivity) SetType(v ActivityType)`

SetType sets Type field to given value.

### HasType

`func (o *SummaryActivity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSportType

`func (o *SummaryActivity) GetSportType() SportType`

GetSportType returns the SportType field if non-nil, zero value otherwise.

### GetSportTypeOk

`func (o *SummaryActivity) GetSportTypeOk() (*SportType, bool)`

GetSportTypeOk returns a tuple with the SportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSportType

`func (o *SummaryActivity) SetSportType(v SportType)`

SetSportType sets SportType field to given value.

### HasSportType

`func (o *SummaryActivity) HasSportType() bool`

HasSportType returns a boolean if a field has been set.

### GetStartDate

`func (o *SummaryActivity) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SummaryActivity) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SummaryActivity) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SummaryActivity) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartDateLocal

`func (o *SummaryActivity) GetStartDateLocal() time.Time`

GetStartDateLocal returns the StartDateLocal field if non-nil, zero value otherwise.

### GetStartDateLocalOk

`func (o *SummaryActivity) GetStartDateLocalOk() (*time.Time, bool)`

GetStartDateLocalOk returns a tuple with the StartDateLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateLocal

`func (o *SummaryActivity) SetStartDateLocal(v time.Time)`

SetStartDateLocal sets StartDateLocal field to given value.

### HasStartDateLocal

`func (o *SummaryActivity) HasStartDateLocal() bool`

HasStartDateLocal returns a boolean if a field has been set.

### GetTimezone

`func (o *SummaryActivity) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SummaryActivity) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SummaryActivity) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SummaryActivity) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartLatlng

`func (o *SummaryActivity) GetStartLatlng() []float32`

GetStartLatlng returns the StartLatlng field if non-nil, zero value otherwise.

### GetStartLatlngOk

`func (o *SummaryActivity) GetStartLatlngOk() (*[]float32, bool)`

GetStartLatlngOk returns a tuple with the StartLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLatlng

`func (o *SummaryActivity) SetStartLatlng(v []float32)`

SetStartLatlng sets StartLatlng field to given value.

### HasStartLatlng

`func (o *SummaryActivity) HasStartLatlng() bool`

HasStartLatlng returns a boolean if a field has been set.

### GetEndLatlng

`func (o *SummaryActivity) GetEndLatlng() []float32`

GetEndLatlng returns the EndLatlng field if non-nil, zero value otherwise.

### GetEndLatlngOk

`func (o *SummaryActivity) GetEndLatlngOk() (*[]float32, bool)`

GetEndLatlngOk returns a tuple with the EndLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLatlng

`func (o *SummaryActivity) SetEndLatlng(v []float32)`

SetEndLatlng sets EndLatlng field to given value.

### HasEndLatlng

`func (o *SummaryActivity) HasEndLatlng() bool`

HasEndLatlng returns a boolean if a field has been set.

### GetAchievementCount

`func (o *SummaryActivity) GetAchievementCount() int32`

GetAchievementCount returns the AchievementCount field if non-nil, zero value otherwise.

### GetAchievementCountOk

`func (o *SummaryActivity) GetAchievementCountOk() (*int32, bool)`

GetAchievementCountOk returns a tuple with the AchievementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementCount

`func (o *SummaryActivity) SetAchievementCount(v int32)`

SetAchievementCount sets AchievementCount field to given value.

### HasAchievementCount

`func (o *SummaryActivity) HasAchievementCount() bool`

HasAchievementCount returns a boolean if a field has been set.

### GetKudosCount

`func (o *SummaryActivity) GetKudosCount() int32`

GetKudosCount returns the KudosCount field if non-nil, zero value otherwise.

### GetKudosCountOk

`func (o *SummaryActivity) GetKudosCountOk() (*int32, bool)`

GetKudosCountOk returns a tuple with the KudosCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKudosCount

`func (o *SummaryActivity) SetKudosCount(v int32)`

SetKudosCount sets KudosCount field to given value.

### HasKudosCount

`func (o *SummaryActivity) HasKudosCount() bool`

HasKudosCount returns a boolean if a field has been set.

### GetCommentCount

`func (o *SummaryActivity) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *SummaryActivity) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *SummaryActivity) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *SummaryActivity) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetAthleteCount

`func (o *SummaryActivity) GetAthleteCount() int32`

GetAthleteCount returns the AthleteCount field if non-nil, zero value otherwise.

### GetAthleteCountOk

`func (o *SummaryActivity) GetAthleteCountOk() (*int32, bool)`

GetAthleteCountOk returns a tuple with the AthleteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthleteCount

`func (o *SummaryActivity) SetAthleteCount(v int32)`

SetAthleteCount sets AthleteCount field to given value.

### HasAthleteCount

`func (o *SummaryActivity) HasAthleteCount() bool`

HasAthleteCount returns a boolean if a field has been set.

### GetPhotoCount

`func (o *SummaryActivity) GetPhotoCount() int32`

GetPhotoCount returns the PhotoCount field if non-nil, zero value otherwise.

### GetPhotoCountOk

`func (o *SummaryActivity) GetPhotoCountOk() (*int32, bool)`

GetPhotoCountOk returns a tuple with the PhotoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoCount

`func (o *SummaryActivity) SetPhotoCount(v int32)`

SetPhotoCount sets PhotoCount field to given value.

### HasPhotoCount

`func (o *SummaryActivity) HasPhotoCount() bool`

HasPhotoCount returns a boolean if a field has been set.

### GetTotalPhotoCount

`func (o *SummaryActivity) GetTotalPhotoCount() int32`

GetTotalPhotoCount returns the TotalPhotoCount field if non-nil, zero value otherwise.

### GetTotalPhotoCountOk

`func (o *SummaryActivity) GetTotalPhotoCountOk() (*int32, bool)`

GetTotalPhotoCountOk returns a tuple with the TotalPhotoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhotoCount

`func (o *SummaryActivity) SetTotalPhotoCount(v int32)`

SetTotalPhotoCount sets TotalPhotoCount field to given value.

### HasTotalPhotoCount

`func (o *SummaryActivity) HasTotalPhotoCount() bool`

HasTotalPhotoCount returns a boolean if a field has been set.

### GetMap

`func (o *SummaryActivity) GetMap() PolylineMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *SummaryActivity) GetMapOk() (*PolylineMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *SummaryActivity) SetMap(v PolylineMap)`

SetMap sets Map field to given value.

### HasMap

`func (o *SummaryActivity) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetDeviceName

`func (o *SummaryActivity) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SummaryActivity) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SummaryActivity) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SummaryActivity) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetTrainer

`func (o *SummaryActivity) GetTrainer() bool`

GetTrainer returns the Trainer field if non-nil, zero value otherwise.

### GetTrainerOk

`func (o *SummaryActivity) GetTrainerOk() (*bool, bool)`

GetTrainerOk returns a tuple with the Trainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainer

`func (o *SummaryActivity) SetTrainer(v bool)`

SetTrainer sets Trainer field to given value.

### HasTrainer

`func (o *SummaryActivity) HasTrainer() bool`

HasTrainer returns a boolean if a field has been set.

### GetCommute

`func (o *SummaryActivity) GetCommute() bool`

GetCommute returns the Commute field if non-nil, zero value otherwise.

### GetCommuteOk

`func (o *SummaryActivity) GetCommuteOk() (*bool, bool)`

GetCommuteOk returns a tuple with the Commute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommute

`func (o *SummaryActivity) SetCommute(v bool)`

SetCommute sets Commute field to given value.

### HasCommute

`func (o *SummaryActivity) HasCommute() bool`

HasCommute returns a boolean if a field has been set.

### GetManual

`func (o *SummaryActivity) GetManual() bool`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *SummaryActivity) GetManualOk() (*bool, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *SummaryActivity) SetManual(v bool)`

SetManual sets Manual field to given value.

### HasManual

`func (o *SummaryActivity) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetPrivate

`func (o *SummaryActivity) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *SummaryActivity) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *SummaryActivity) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *SummaryActivity) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetFlagged

`func (o *SummaryActivity) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *SummaryActivity) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *SummaryActivity) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *SummaryActivity) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetWorkoutType

`func (o *SummaryActivity) GetWorkoutType() int32`

GetWorkoutType returns the WorkoutType field if non-nil, zero value otherwise.

### GetWorkoutTypeOk

`func (o *SummaryActivity) GetWorkoutTypeOk() (*int32, bool)`

GetWorkoutTypeOk returns a tuple with the WorkoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkoutType

`func (o *SummaryActivity) SetWorkoutType(v int32)`

SetWorkoutType sets WorkoutType field to given value.

### HasWorkoutType

`func (o *SummaryActivity) HasWorkoutType() bool`

HasWorkoutType returns a boolean if a field has been set.

### GetUploadIdStr

`func (o *SummaryActivity) GetUploadIdStr() string`

GetUploadIdStr returns the UploadIdStr field if non-nil, zero value otherwise.

### GetUploadIdStrOk

`func (o *SummaryActivity) GetUploadIdStrOk() (*string, bool)`

GetUploadIdStrOk returns a tuple with the UploadIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadIdStr

`func (o *SummaryActivity) SetUploadIdStr(v string)`

SetUploadIdStr sets UploadIdStr field to given value.

### HasUploadIdStr

`func (o *SummaryActivity) HasUploadIdStr() bool`

HasUploadIdStr returns a boolean if a field has been set.

### GetAverageSpeed

`func (o *SummaryActivity) GetAverageSpeed() float32`

GetAverageSpeed returns the AverageSpeed field if non-nil, zero value otherwise.

### GetAverageSpeedOk

`func (o *SummaryActivity) GetAverageSpeedOk() (*float32, bool)`

GetAverageSpeedOk returns a tuple with the AverageSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageSpeed

`func (o *SummaryActivity) SetAverageSpeed(v float32)`

SetAverageSpeed sets AverageSpeed field to given value.

### HasAverageSpeed

`func (o *SummaryActivity) HasAverageSpeed() bool`

HasAverageSpeed returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *SummaryActivity) GetMaxSpeed() float32`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *SummaryActivity) GetMaxSpeedOk() (*float32, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *SummaryActivity) SetMaxSpeed(v float32)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *SummaryActivity) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetHasKudoed

`func (o *SummaryActivity) GetHasKudoed() bool`

GetHasKudoed returns the HasKudoed field if non-nil, zero value otherwise.

### GetHasKudoedOk

`func (o *SummaryActivity) GetHasKudoedOk() (*bool, bool)`

GetHasKudoedOk returns a tuple with the HasKudoed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKudoed

`func (o *SummaryActivity) SetHasKudoed(v bool)`

SetHasKudoed sets HasKudoed field to given value.

### HasHasKudoed

`func (o *SummaryActivity) HasHasKudoed() bool`

HasHasKudoed returns a boolean if a field has been set.

### GetHideFromHome

`func (o *SummaryActivity) GetHideFromHome() bool`

GetHideFromHome returns the HideFromHome field if non-nil, zero value otherwise.

### GetHideFromHomeOk

`func (o *SummaryActivity) GetHideFromHomeOk() (*bool, bool)`

GetHideFromHomeOk returns a tuple with the HideFromHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromHome

`func (o *SummaryActivity) SetHideFromHome(v bool)`

SetHideFromHome sets HideFromHome field to given value.

### HasHideFromHome

`func (o *SummaryActivity) HasHideFromHome() bool`

HasHideFromHome returns a boolean if a field has been set.

### GetGearId

`func (o *SummaryActivity) GetGearId() string`

GetGearId returns the GearId field if non-nil, zero value otherwise.

### GetGearIdOk

`func (o *SummaryActivity) GetGearIdOk() (*string, bool)`

GetGearIdOk returns a tuple with the GearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGearId

`func (o *SummaryActivity) SetGearId(v string)`

SetGearId sets GearId field to given value.

### HasGearId

`func (o *SummaryActivity) HasGearId() bool`

HasGearId returns a boolean if a field has been set.

### GetKilojoules

`func (o *SummaryActivity) GetKilojoules() float32`

GetKilojoules returns the Kilojoules field if non-nil, zero value otherwise.

### GetKilojoulesOk

`func (o *SummaryActivity) GetKilojoulesOk() (*float32, bool)`

GetKilojoulesOk returns a tuple with the Kilojoules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKilojoules

`func (o *SummaryActivity) SetKilojoules(v float32)`

SetKilojoules sets Kilojoules field to given value.

### HasKilojoules

`func (o *SummaryActivity) HasKilojoules() bool`

HasKilojoules returns a boolean if a field has been set.

### GetAverageWatts

`func (o *SummaryActivity) GetAverageWatts() float32`

GetAverageWatts returns the AverageWatts field if non-nil, zero value otherwise.

### GetAverageWattsOk

`func (o *SummaryActivity) GetAverageWattsOk() (*float32, bool)`

GetAverageWattsOk returns a tuple with the AverageWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWatts

`func (o *SummaryActivity) SetAverageWatts(v float32)`

SetAverageWatts sets AverageWatts field to given value.

### HasAverageWatts

`func (o *SummaryActivity) HasAverageWatts() bool`

HasAverageWatts returns a boolean if a field has been set.

### GetDeviceWatts

`func (o *SummaryActivity) GetDeviceWatts() bool`

GetDeviceWatts returns the DeviceWatts field if non-nil, zero value otherwise.

### GetDeviceWattsOk

`func (o *SummaryActivity) GetDeviceWattsOk() (*bool, bool)`

GetDeviceWattsOk returns a tuple with the DeviceWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceWatts

`func (o *SummaryActivity) SetDeviceWatts(v bool)`

SetDeviceWatts sets DeviceWatts field to given value.

### HasDeviceWatts

`func (o *SummaryActivity) HasDeviceWatts() bool`

HasDeviceWatts returns a boolean if a field has been set.

### GetMaxWatts

`func (o *SummaryActivity) GetMaxWatts() int32`

GetMaxWatts returns the MaxWatts field if non-nil, zero value otherwise.

### GetMaxWattsOk

`func (o *SummaryActivity) GetMaxWattsOk() (*int32, bool)`

GetMaxWattsOk returns a tuple with the MaxWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWatts

`func (o *SummaryActivity) SetMaxWatts(v int32)`

SetMaxWatts sets MaxWatts field to given value.

### HasMaxWatts

`func (o *SummaryActivity) HasMaxWatts() bool`

HasMaxWatts returns a boolean if a field has been set.

### GetWeightedAverageWatts

`func (o *SummaryActivity) GetWeightedAverageWatts() int32`

GetWeightedAverageWatts returns the WeightedAverageWatts field if non-nil, zero value otherwise.

### GetWeightedAverageWattsOk

`func (o *SummaryActivity) GetWeightedAverageWattsOk() (*int32, bool)`

GetWeightedAverageWattsOk returns a tuple with the WeightedAverageWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAverageWatts

`func (o *SummaryActivity) SetWeightedAverageWatts(v int32)`

SetWeightedAverageWatts sets WeightedAverageWatts field to given value.

### HasWeightedAverageWatts

`func (o *SummaryActivity) HasWeightedAverageWatts() bool`

HasWeightedAverageWatts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


