// Code generated by generator, DO NOT EDIT.
package cloudtrail_input

type zone struct {
	name   *string
	offset *int
	isDST  *bool
}
type zoneTrans struct {
	when  *int64
	index *uint8
	isstd *bool
	isutc *bool
}
type Location struct {
	name       *string
	zone       []zone
	tx         []zoneTrans
	extend     *string
	cacheStart *int64
	cacheEnd   *int64
	cacheZone  *zone
}
type Time struct {
	wall *uint64
	ext  *int64
	loc  *Location
}
type EventCategory string
type LookupAttributeKey string
type LookupAttribute struct {
	AttributeKey   *LookupAttributeKey
	AttributeValue *string
}
type LookupEventsInput struct {
	EndTime          *Time
	EventCategory    *EventCategory
	LookupAttributes []LookupAttribute
	MaxResults       *int32
	NextToken        *string
	StartTime        *Time
}
