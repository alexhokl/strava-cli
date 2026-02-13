# DetailedActivity

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
**Description** | Pointer to **string** | The description of the activity | [optional] 
**Photos** | Pointer to [**PhotosSummary**](PhotosSummary.md) |  | [optional] 
**Gear** | Pointer to [**SummaryGear**](SummaryGear.md) |  | [optional] 
**Calories** | Pointer to **float32** | The number of kilocalories consumed during this activity | [optional] 
**SegmentEfforts** | Pointer to [**[]DetailedSegmentEffort**](DetailedSegmentEffort.md) |  | [optional] 
**EmbedToken** | Pointer to **string** | The token used to embed a Strava activity | [optional] 
**SplitsMetric** | Pointer to [**[]Split**](Split.md) | The splits of this activity in metric units (for runs) | [optional] 
**SplitsStandard** | Pointer to [**[]Split**](Split.md) | The splits of this activity in imperial units (for runs) | [optional] 
**Laps** | Pointer to [**[]Lap**](Lap.md) |  | [optional] 
**BestEfforts** | Pointer to [**[]DetailedSegmentEffort**](DetailedSegmentEffort.md) |  | [optional] 

## Methods

### NewDetailedActivity

`func NewDetailedActivity() *DetailedActivity`

NewDetailedActivity instantiates a new DetailedActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedActivityWithDefaults

`func NewDetailedActivityWithDefaults() *DetailedActivity`

NewDetailedActivityWithDefaults instantiates a new DetailedActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DetailedActivity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedActivity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedActivity) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedActivity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *DetailedActivity) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *DetailedActivity) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *DetailedActivity) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *DetailedActivity) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUploadId

`func (o *DetailedActivity) GetUploadId() int64`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *DetailedActivity) GetUploadIdOk() (*int64, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *DetailedActivity) SetUploadId(v int64)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *DetailedActivity) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetAthlete

`func (o *DetailedActivity) GetAthlete() MetaAthlete`

GetAthlete returns the Athlete field if non-nil, zero value otherwise.

### GetAthleteOk

`func (o *DetailedActivity) GetAthleteOk() (*MetaAthlete, bool)`

GetAthleteOk returns a tuple with the Athlete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthlete

`func (o *DetailedActivity) SetAthlete(v MetaAthlete)`

SetAthlete sets Athlete field to given value.

### HasAthlete

`func (o *DetailedActivity) HasAthlete() bool`

HasAthlete returns a boolean if a field has been set.

### GetName

`func (o *DetailedActivity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedActivity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedActivity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedActivity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDistance

`func (o *DetailedActivity) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *DetailedActivity) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *DetailedActivity) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *DetailedActivity) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetMovingTime

`func (o *DetailedActivity) GetMovingTime() int32`

GetMovingTime returns the MovingTime field if non-nil, zero value otherwise.

### GetMovingTimeOk

`func (o *DetailedActivity) GetMovingTimeOk() (*int32, bool)`

GetMovingTimeOk returns a tuple with the MovingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovingTime

`func (o *DetailedActivity) SetMovingTime(v int32)`

SetMovingTime sets MovingTime field to given value.

### HasMovingTime

`func (o *DetailedActivity) HasMovingTime() bool`

HasMovingTime returns a boolean if a field has been set.

### GetElapsedTime

`func (o *DetailedActivity) GetElapsedTime() int32`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *DetailedActivity) GetElapsedTimeOk() (*int32, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *DetailedActivity) SetElapsedTime(v int32)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *DetailedActivity) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetTotalElevationGain

`func (o *DetailedActivity) GetTotalElevationGain() float32`

GetTotalElevationGain returns the TotalElevationGain field if non-nil, zero value otherwise.

### GetTotalElevationGainOk

`func (o *DetailedActivity) GetTotalElevationGainOk() (*float32, bool)`

GetTotalElevationGainOk returns a tuple with the TotalElevationGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalElevationGain

`func (o *DetailedActivity) SetTotalElevationGain(v float32)`

SetTotalElevationGain sets TotalElevationGain field to given value.

### HasTotalElevationGain

`func (o *DetailedActivity) HasTotalElevationGain() bool`

HasTotalElevationGain returns a boolean if a field has been set.

### GetElevHigh

`func (o *DetailedActivity) GetElevHigh() float32`

GetElevHigh returns the ElevHigh field if non-nil, zero value otherwise.

### GetElevHighOk

`func (o *DetailedActivity) GetElevHighOk() (*float32, bool)`

GetElevHighOk returns a tuple with the ElevHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevHigh

`func (o *DetailedActivity) SetElevHigh(v float32)`

SetElevHigh sets ElevHigh field to given value.

### HasElevHigh

`func (o *DetailedActivity) HasElevHigh() bool`

HasElevHigh returns a boolean if a field has been set.

### GetElevLow

`func (o *DetailedActivity) GetElevLow() float32`

GetElevLow returns the ElevLow field if non-nil, zero value otherwise.

### GetElevLowOk

`func (o *DetailedActivity) GetElevLowOk() (*float32, bool)`

GetElevLowOk returns a tuple with the ElevLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevLow

`func (o *DetailedActivity) SetElevLow(v float32)`

SetElevLow sets ElevLow field to given value.

### HasElevLow

`func (o *DetailedActivity) HasElevLow() bool`

HasElevLow returns a boolean if a field has been set.

### GetType

`func (o *DetailedActivity) GetType() ActivityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DetailedActivity) GetTypeOk() (*ActivityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DetailedActivity) SetType(v ActivityType)`

SetType sets Type field to given value.

### HasType

`func (o *DetailedActivity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSportType

`func (o *DetailedActivity) GetSportType() SportType`

GetSportType returns the SportType field if non-nil, zero value otherwise.

### GetSportTypeOk

`func (o *DetailedActivity) GetSportTypeOk() (*SportType, bool)`

GetSportTypeOk returns a tuple with the SportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSportType

`func (o *DetailedActivity) SetSportType(v SportType)`

SetSportType sets SportType field to given value.

### HasSportType

`func (o *DetailedActivity) HasSportType() bool`

HasSportType returns a boolean if a field has been set.

### GetStartDate

`func (o *DetailedActivity) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DetailedActivity) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DetailedActivity) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DetailedActivity) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartDateLocal

`func (o *DetailedActivity) GetStartDateLocal() time.Time`

GetStartDateLocal returns the StartDateLocal field if non-nil, zero value otherwise.

### GetStartDateLocalOk

`func (o *DetailedActivity) GetStartDateLocalOk() (*time.Time, bool)`

GetStartDateLocalOk returns a tuple with the StartDateLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateLocal

`func (o *DetailedActivity) SetStartDateLocal(v time.Time)`

SetStartDateLocal sets StartDateLocal field to given value.

### HasStartDateLocal

`func (o *DetailedActivity) HasStartDateLocal() bool`

HasStartDateLocal returns a boolean if a field has been set.

### GetTimezone

`func (o *DetailedActivity) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DetailedActivity) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DetailedActivity) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *DetailedActivity) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStartLatlng

`func (o *DetailedActivity) GetStartLatlng() []float32`

GetStartLatlng returns the StartLatlng field if non-nil, zero value otherwise.

### GetStartLatlngOk

`func (o *DetailedActivity) GetStartLatlngOk() (*[]float32, bool)`

GetStartLatlngOk returns a tuple with the StartLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLatlng

`func (o *DetailedActivity) SetStartLatlng(v []float32)`

SetStartLatlng sets StartLatlng field to given value.

### HasStartLatlng

`func (o *DetailedActivity) HasStartLatlng() bool`

HasStartLatlng returns a boolean if a field has been set.

### GetEndLatlng

`func (o *DetailedActivity) GetEndLatlng() []float32`

GetEndLatlng returns the EndLatlng field if non-nil, zero value otherwise.

### GetEndLatlngOk

`func (o *DetailedActivity) GetEndLatlngOk() (*[]float32, bool)`

GetEndLatlngOk returns a tuple with the EndLatlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLatlng

`func (o *DetailedActivity) SetEndLatlng(v []float32)`

SetEndLatlng sets EndLatlng field to given value.

### HasEndLatlng

`func (o *DetailedActivity) HasEndLatlng() bool`

HasEndLatlng returns a boolean if a field has been set.

### GetAchievementCount

`func (o *DetailedActivity) GetAchievementCount() int32`

GetAchievementCount returns the AchievementCount field if non-nil, zero value otherwise.

### GetAchievementCountOk

`func (o *DetailedActivity) GetAchievementCountOk() (*int32, bool)`

GetAchievementCountOk returns a tuple with the AchievementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementCount

`func (o *DetailedActivity) SetAchievementCount(v int32)`

SetAchievementCount sets AchievementCount field to given value.

### HasAchievementCount

`func (o *DetailedActivity) HasAchievementCount() bool`

HasAchievementCount returns a boolean if a field has been set.

### GetKudosCount

`func (o *DetailedActivity) GetKudosCount() int32`

GetKudosCount returns the KudosCount field if non-nil, zero value otherwise.

### GetKudosCountOk

`func (o *DetailedActivity) GetKudosCountOk() (*int32, bool)`

GetKudosCountOk returns a tuple with the KudosCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKudosCount

`func (o *DetailedActivity) SetKudosCount(v int32)`

SetKudosCount sets KudosCount field to given value.

### HasKudosCount

`func (o *DetailedActivity) HasKudosCount() bool`

HasKudosCount returns a boolean if a field has been set.

### GetCommentCount

`func (o *DetailedActivity) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *DetailedActivity) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *DetailedActivity) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *DetailedActivity) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetAthleteCount

`func (o *DetailedActivity) GetAthleteCount() int32`

GetAthleteCount returns the AthleteCount field if non-nil, zero value otherwise.

### GetAthleteCountOk

`func (o *DetailedActivity) GetAthleteCountOk() (*int32, bool)`

GetAthleteCountOk returns a tuple with the AthleteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthleteCount

`func (o *DetailedActivity) SetAthleteCount(v int32)`

SetAthleteCount sets AthleteCount field to given value.

### HasAthleteCount

`func (o *DetailedActivity) HasAthleteCount() bool`

HasAthleteCount returns a boolean if a field has been set.

### GetPhotoCount

`func (o *DetailedActivity) GetPhotoCount() int32`

GetPhotoCount returns the PhotoCount field if non-nil, zero value otherwise.

### GetPhotoCountOk

`func (o *DetailedActivity) GetPhotoCountOk() (*int32, bool)`

GetPhotoCountOk returns a tuple with the PhotoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoCount

`func (o *DetailedActivity) SetPhotoCount(v int32)`

SetPhotoCount sets PhotoCount field to given value.

### HasPhotoCount

`func (o *DetailedActivity) HasPhotoCount() bool`

HasPhotoCount returns a boolean if a field has been set.

### GetTotalPhotoCount

`func (o *DetailedActivity) GetTotalPhotoCount() int32`

GetTotalPhotoCount returns the TotalPhotoCount field if non-nil, zero value otherwise.

### GetTotalPhotoCountOk

`func (o *DetailedActivity) GetTotalPhotoCountOk() (*int32, bool)`

GetTotalPhotoCountOk returns a tuple with the TotalPhotoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhotoCount

`func (o *DetailedActivity) SetTotalPhotoCount(v int32)`

SetTotalPhotoCount sets TotalPhotoCount field to given value.

### HasTotalPhotoCount

`func (o *DetailedActivity) HasTotalPhotoCount() bool`

HasTotalPhotoCount returns a boolean if a field has been set.

### GetMap

`func (o *DetailedActivity) GetMap() PolylineMap`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *DetailedActivity) GetMapOk() (*PolylineMap, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *DetailedActivity) SetMap(v PolylineMap)`

SetMap sets Map field to given value.

### HasMap

`func (o *DetailedActivity) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetDeviceName

`func (o *DetailedActivity) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DetailedActivity) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DetailedActivity) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DetailedActivity) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetTrainer

`func (o *DetailedActivity) GetTrainer() bool`

GetTrainer returns the Trainer field if non-nil, zero value otherwise.

### GetTrainerOk

`func (o *DetailedActivity) GetTrainerOk() (*bool, bool)`

GetTrainerOk returns a tuple with the Trainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainer

`func (o *DetailedActivity) SetTrainer(v bool)`

SetTrainer sets Trainer field to given value.

### HasTrainer

`func (o *DetailedActivity) HasTrainer() bool`

HasTrainer returns a boolean if a field has been set.

### GetCommute

`func (o *DetailedActivity) GetCommute() bool`

GetCommute returns the Commute field if non-nil, zero value otherwise.

### GetCommuteOk

`func (o *DetailedActivity) GetCommuteOk() (*bool, bool)`

GetCommuteOk returns a tuple with the Commute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommute

`func (o *DetailedActivity) SetCommute(v bool)`

SetCommute sets Commute field to given value.

### HasCommute

`func (o *DetailedActivity) HasCommute() bool`

HasCommute returns a boolean if a field has been set.

### GetManual

`func (o *DetailedActivity) GetManual() bool`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *DetailedActivity) GetManualOk() (*bool, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *DetailedActivity) SetManual(v bool)`

SetManual sets Manual field to given value.

### HasManual

`func (o *DetailedActivity) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetPrivate

`func (o *DetailedActivity) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *DetailedActivity) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *DetailedActivity) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *DetailedActivity) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetFlagged

`func (o *DetailedActivity) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *DetailedActivity) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *DetailedActivity) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.

### HasFlagged

`func (o *DetailedActivity) HasFlagged() bool`

HasFlagged returns a boolean if a field has been set.

### GetWorkoutType

`func (o *DetailedActivity) GetWorkoutType() int32`

GetWorkoutType returns the WorkoutType field if non-nil, zero value otherwise.

### GetWorkoutTypeOk

`func (o *DetailedActivity) GetWorkoutTypeOk() (*int32, bool)`

GetWorkoutTypeOk returns a tuple with the WorkoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkoutType

`func (o *DetailedActivity) SetWorkoutType(v int32)`

SetWorkoutType sets WorkoutType field to given value.

### HasWorkoutType

`func (o *DetailedActivity) HasWorkoutType() bool`

HasWorkoutType returns a boolean if a field has been set.

### GetUploadIdStr

`func (o *DetailedActivity) GetUploadIdStr() string`

GetUploadIdStr returns the UploadIdStr field if non-nil, zero value otherwise.

### GetUploadIdStrOk

`func (o *DetailedActivity) GetUploadIdStrOk() (*string, bool)`

GetUploadIdStrOk returns a tuple with the UploadIdStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadIdStr

`func (o *DetailedActivity) SetUploadIdStr(v string)`

SetUploadIdStr sets UploadIdStr field to given value.

### HasUploadIdStr

`func (o *DetailedActivity) HasUploadIdStr() bool`

HasUploadIdStr returns a boolean if a field has been set.

### GetAverageSpeed

`func (o *DetailedActivity) GetAverageSpeed() float32`

GetAverageSpeed returns the AverageSpeed field if non-nil, zero value otherwise.

### GetAverageSpeedOk

`func (o *DetailedActivity) GetAverageSpeedOk() (*float32, bool)`

GetAverageSpeedOk returns a tuple with the AverageSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageSpeed

`func (o *DetailedActivity) SetAverageSpeed(v float32)`

SetAverageSpeed sets AverageSpeed field to given value.

### HasAverageSpeed

`func (o *DetailedActivity) HasAverageSpeed() bool`

HasAverageSpeed returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *DetailedActivity) GetMaxSpeed() float32`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *DetailedActivity) GetMaxSpeedOk() (*float32, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *DetailedActivity) SetMaxSpeed(v float32)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *DetailedActivity) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetHasKudoed

`func (o *DetailedActivity) GetHasKudoed() bool`

GetHasKudoed returns the HasKudoed field if non-nil, zero value otherwise.

### GetHasKudoedOk

`func (o *DetailedActivity) GetHasKudoedOk() (*bool, bool)`

GetHasKudoedOk returns a tuple with the HasKudoed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKudoed

`func (o *DetailedActivity) SetHasKudoed(v bool)`

SetHasKudoed sets HasKudoed field to given value.

### HasHasKudoed

`func (o *DetailedActivity) HasHasKudoed() bool`

HasHasKudoed returns a boolean if a field has been set.

### GetHideFromHome

`func (o *DetailedActivity) GetHideFromHome() bool`

GetHideFromHome returns the HideFromHome field if non-nil, zero value otherwise.

### GetHideFromHomeOk

`func (o *DetailedActivity) GetHideFromHomeOk() (*bool, bool)`

GetHideFromHomeOk returns a tuple with the HideFromHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromHome

`func (o *DetailedActivity) SetHideFromHome(v bool)`

SetHideFromHome sets HideFromHome field to given value.

### HasHideFromHome

`func (o *DetailedActivity) HasHideFromHome() bool`

HasHideFromHome returns a boolean if a field has been set.

### GetGearId

`func (o *DetailedActivity) GetGearId() string`

GetGearId returns the GearId field if non-nil, zero value otherwise.

### GetGearIdOk

`func (o *DetailedActivity) GetGearIdOk() (*string, bool)`

GetGearIdOk returns a tuple with the GearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGearId

`func (o *DetailedActivity) SetGearId(v string)`

SetGearId sets GearId field to given value.

### HasGearId

`func (o *DetailedActivity) HasGearId() bool`

HasGearId returns a boolean if a field has been set.

### GetKilojoules

`func (o *DetailedActivity) GetKilojoules() float32`

GetKilojoules returns the Kilojoules field if non-nil, zero value otherwise.

### GetKilojoulesOk

`func (o *DetailedActivity) GetKilojoulesOk() (*float32, bool)`

GetKilojoulesOk returns a tuple with the Kilojoules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKilojoules

`func (o *DetailedActivity) SetKilojoules(v float32)`

SetKilojoules sets Kilojoules field to given value.

### HasKilojoules

`func (o *DetailedActivity) HasKilojoules() bool`

HasKilojoules returns a boolean if a field has been set.

### GetAverageWatts

`func (o *DetailedActivity) GetAverageWatts() float32`

GetAverageWatts returns the AverageWatts field if non-nil, zero value otherwise.

### GetAverageWattsOk

`func (o *DetailedActivity) GetAverageWattsOk() (*float32, bool)`

GetAverageWattsOk returns a tuple with the AverageWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWatts

`func (o *DetailedActivity) SetAverageWatts(v float32)`

SetAverageWatts sets AverageWatts field to given value.

### HasAverageWatts

`func (o *DetailedActivity) HasAverageWatts() bool`

HasAverageWatts returns a boolean if a field has been set.

### GetDeviceWatts

`func (o *DetailedActivity) GetDeviceWatts() bool`

GetDeviceWatts returns the DeviceWatts field if non-nil, zero value otherwise.

### GetDeviceWattsOk

`func (o *DetailedActivity) GetDeviceWattsOk() (*bool, bool)`

GetDeviceWattsOk returns a tuple with the DeviceWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceWatts

`func (o *DetailedActivity) SetDeviceWatts(v bool)`

SetDeviceWatts sets DeviceWatts field to given value.

### HasDeviceWatts

`func (o *DetailedActivity) HasDeviceWatts() bool`

HasDeviceWatts returns a boolean if a field has been set.

### GetMaxWatts

`func (o *DetailedActivity) GetMaxWatts() int32`

GetMaxWatts returns the MaxWatts field if non-nil, zero value otherwise.

### GetMaxWattsOk

`func (o *DetailedActivity) GetMaxWattsOk() (*int32, bool)`

GetMaxWattsOk returns a tuple with the MaxWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWatts

`func (o *DetailedActivity) SetMaxWatts(v int32)`

SetMaxWatts sets MaxWatts field to given value.

### HasMaxWatts

`func (o *DetailedActivity) HasMaxWatts() bool`

HasMaxWatts returns a boolean if a field has been set.

### GetWeightedAverageWatts

`func (o *DetailedActivity) GetWeightedAverageWatts() int32`

GetWeightedAverageWatts returns the WeightedAverageWatts field if non-nil, zero value otherwise.

### GetWeightedAverageWattsOk

`func (o *DetailedActivity) GetWeightedAverageWattsOk() (*int32, bool)`

GetWeightedAverageWattsOk returns a tuple with the WeightedAverageWatts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAverageWatts

`func (o *DetailedActivity) SetWeightedAverageWatts(v int32)`

SetWeightedAverageWatts sets WeightedAverageWatts field to given value.

### HasWeightedAverageWatts

`func (o *DetailedActivity) HasWeightedAverageWatts() bool`

HasWeightedAverageWatts returns a boolean if a field has been set.

### GetDescription

`func (o *DetailedActivity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DetailedActivity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DetailedActivity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DetailedActivity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPhotos

`func (o *DetailedActivity) GetPhotos() PhotosSummary`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *DetailedActivity) GetPhotosOk() (*PhotosSummary, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *DetailedActivity) SetPhotos(v PhotosSummary)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *DetailedActivity) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### GetGear

`func (o *DetailedActivity) GetGear() SummaryGear`

GetGear returns the Gear field if non-nil, zero value otherwise.

### GetGearOk

`func (o *DetailedActivity) GetGearOk() (*SummaryGear, bool)`

GetGearOk returns a tuple with the Gear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGear

`func (o *DetailedActivity) SetGear(v SummaryGear)`

SetGear sets Gear field to given value.

### HasGear

`func (o *DetailedActivity) HasGear() bool`

HasGear returns a boolean if a field has been set.

### GetCalories

`func (o *DetailedActivity) GetCalories() float32`

GetCalories returns the Calories field if non-nil, zero value otherwise.

### GetCaloriesOk

`func (o *DetailedActivity) GetCaloriesOk() (*float32, bool)`

GetCaloriesOk returns a tuple with the Calories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalories

`func (o *DetailedActivity) SetCalories(v float32)`

SetCalories sets Calories field to given value.

### HasCalories

`func (o *DetailedActivity) HasCalories() bool`

HasCalories returns a boolean if a field has been set.

### GetSegmentEfforts

`func (o *DetailedActivity) GetSegmentEfforts() []DetailedSegmentEffort`

GetSegmentEfforts returns the SegmentEfforts field if non-nil, zero value otherwise.

### GetSegmentEffortsOk

`func (o *DetailedActivity) GetSegmentEffortsOk() (*[]DetailedSegmentEffort, bool)`

GetSegmentEffortsOk returns a tuple with the SegmentEfforts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentEfforts

`func (o *DetailedActivity) SetSegmentEfforts(v []DetailedSegmentEffort)`

SetSegmentEfforts sets SegmentEfforts field to given value.

### HasSegmentEfforts

`func (o *DetailedActivity) HasSegmentEfforts() bool`

HasSegmentEfforts returns a boolean if a field has been set.

### GetEmbedToken

`func (o *DetailedActivity) GetEmbedToken() string`

GetEmbedToken returns the EmbedToken field if non-nil, zero value otherwise.

### GetEmbedTokenOk

`func (o *DetailedActivity) GetEmbedTokenOk() (*string, bool)`

GetEmbedTokenOk returns a tuple with the EmbedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedToken

`func (o *DetailedActivity) SetEmbedToken(v string)`

SetEmbedToken sets EmbedToken field to given value.

### HasEmbedToken

`func (o *DetailedActivity) HasEmbedToken() bool`

HasEmbedToken returns a boolean if a field has been set.

### GetSplitsMetric

`func (o *DetailedActivity) GetSplitsMetric() []Split`

GetSplitsMetric returns the SplitsMetric field if non-nil, zero value otherwise.

### GetSplitsMetricOk

`func (o *DetailedActivity) GetSplitsMetricOk() (*[]Split, bool)`

GetSplitsMetricOk returns a tuple with the SplitsMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitsMetric

`func (o *DetailedActivity) SetSplitsMetric(v []Split)`

SetSplitsMetric sets SplitsMetric field to given value.

### HasSplitsMetric

`func (o *DetailedActivity) HasSplitsMetric() bool`

HasSplitsMetric returns a boolean if a field has been set.

### GetSplitsStandard

`func (o *DetailedActivity) GetSplitsStandard() []Split`

GetSplitsStandard returns the SplitsStandard field if non-nil, zero value otherwise.

### GetSplitsStandardOk

`func (o *DetailedActivity) GetSplitsStandardOk() (*[]Split, bool)`

GetSplitsStandardOk returns a tuple with the SplitsStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitsStandard

`func (o *DetailedActivity) SetSplitsStandard(v []Split)`

SetSplitsStandard sets SplitsStandard field to given value.

### HasSplitsStandard

`func (o *DetailedActivity) HasSplitsStandard() bool`

HasSplitsStandard returns a boolean if a field has been set.

### GetLaps

`func (o *DetailedActivity) GetLaps() []Lap`

GetLaps returns the Laps field if non-nil, zero value otherwise.

### GetLapsOk

`func (o *DetailedActivity) GetLapsOk() (*[]Lap, bool)`

GetLapsOk returns a tuple with the Laps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaps

`func (o *DetailedActivity) SetLaps(v []Lap)`

SetLaps sets Laps field to given value.

### HasLaps

`func (o *DetailedActivity) HasLaps() bool`

HasLaps returns a boolean if a field has been set.

### GetBestEfforts

`func (o *DetailedActivity) GetBestEfforts() []DetailedSegmentEffort`

GetBestEfforts returns the BestEfforts field if non-nil, zero value otherwise.

### GetBestEffortsOk

`func (o *DetailedActivity) GetBestEffortsOk() (*[]DetailedSegmentEffort, bool)`

GetBestEffortsOk returns a tuple with the BestEfforts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestEfforts

`func (o *DetailedActivity) SetBestEfforts(v []DetailedSegmentEffort)`

SetBestEfforts sets BestEfforts field to given value.

### HasBestEfforts

`func (o *DetailedActivity) HasBestEfforts() bool`

HasBestEfforts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


