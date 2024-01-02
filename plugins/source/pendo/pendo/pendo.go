package pendo

import (
	"context"
	"encoding/json"
	"time"
)

type Client interface {
	GetPages(ctx context.Context) ([]Page, error)
	GetFeatures(ctx context.Context) ([]Feature, error)
	GetTrackTypes(ctx context.Context) ([]TrackType, error)
	GetGuides(ctx context.Context) ([]Guide, error)
	GetReports(ctx context.Context) ([]Report, error)
}

type User struct {
	Id          string   `json:"id"`
	Username    string   `json:"username"`
	First       string   `json:"first"`
	Last        string   `json:"last"`
	Role        int      `json:"role"`
	UserType    string   `json:"userType"`
	HasLoggedIn bool     `json:"hasLoggedIn"`
	LastLogin   int64    `json:"lastLogin"`
	VisitorIds  []string `json:"visitorIds"`
}

type Group struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Color             string `json:"color"`
	Length            int    `json:"length"`
	Items             []any  `json:"items"`
	Type              string `json:"type"`
	CreatedByUser     User   `json:"createdByUser,omitempty"`
	CreatedAt         int64  `json:"createdAt,omitempty"`
	LastUpdatedByUser User   `json:"lastUpdatedByUser,omitempty"`
	LastUpdatedAt     int64  `json:"lastUpdatedAt,omitempty"`
}

type Rule struct {
	Rule         string `json:"rule"`
	DesignerHint string `json:"designerHint,omitempty"`
	ParsedRule   string `json:"parsedRule"`
}

type Page struct {
	CreatedByUser     User   `json:"createdByUser"`
	CreatedAt         int64  `json:"createdAt"`
	LastUpdatedByUser User   `json:"lastUpdatedByUser"`
	LastUpdatedAt     int64  `json:"lastUpdatedAt"`
	Kind              string `json:"kind"`
	RootVersionId     string `json:"rootVersionId"`
	StableVersionId   string `json:"stableVersionId"`
	Id                string `json:"id"`
	AppId             int    `json:"appId"`
	Name              string `json:"name"`
	Color             string `json:"color"`
	Group             `json:"group"`
	IsCoreEvent       bool            `json:"isCoreEvent"`
	ValidThrough      int64           `json:"validThrough"`
	Dirty             bool            `json:"dirty"`
	DailyMergeFirst   int             `json:"dailyMergeFirst"`
	DailyRollupFirst  int             `json:"dailyRollupFirst"`
	Rules             []Rule          `json:"rules"`
	Rulesjson         json.RawMessage `json:"rulesjson"`
	IsAutoTagged      bool            `json:"isAutoTagged"`
	SuggestedName     string          `json:"suggestedName,omitempty"`
}

type EventPropertyConfiguration struct {
	Name     string `json:"name"`
	Rule     string `json:"rule"`
	IsActive bool   `json:"isActive"`
	Path     string `json:"path"`
	Selector string `json:"selector"`
	Type     string `json:"type"`
	Pattern  string `json:"pattern"`
}

type Feature struct {
	CreatedByUser               User                         `json:"createdByUser"`
	CreatedAt                   int64                        `json:"createdAt"`
	LastUpdatedByUser           User                         `json:"lastUpdatedByUser"`
	LastUpdatedAt               int64                        `json:"lastUpdatedAt"`
	Kind                        string                       `json:"kind"`
	RootVersionId               string                       `json:"rootVersionId"`
	StableVersionId             string                       `json:"stableVersionId"`
	Id                          string                       `json:"id"`
	AppId                       int                          `json:"appId"`
	Name                        string                       `json:"name"`
	Color                       string                       `json:"color"`
	Group                       Group                        `json:"group"`
	IsCoreEvent                 bool                         `json:"isCoreEvent"`
	ValidThrough                int64                        `json:"validThrough"`
	Dirty                       bool                         `json:"dirty"`
	DailyMergeFirst             int                          `json:"dailyMergeFirst"`
	DailyRollupFirst            int                          `json:"dailyRollupFirst"`
	PageId                      string                       `json:"pageId,omitempty"`
	EventPropertyConfigurations []EventPropertyConfiguration `json:"eventPropertyConfigurations"`
	ElementPathRules            []string                     `json:"elementPathRules"`
	ElementSelectionType        string                       `json:"elementSelectionType,omitempty"`
	SuggestedMatch              string                       `json:"suggestedMatch,omitempty"`
	AppWide                     bool                         `json:"appWide"`
}

type TrackType struct {
	CreatedByUser         User     `json:"createdByUser"`
	CreatedAt             int64    `json:"createdAt"`
	LastUpdatedByUser     User     `json:"lastUpdatedByUser"`
	LastUpdatedAt         int64    `json:"lastUpdatedAt"`
	Kind                  string   `json:"kind"`
	RootVersionId         string   `json:"rootVersionId"`
	StableVersionId       string   `json:"stableVersionId"`
	Id                    string   `json:"id"`
	AppId                 int      `json:"appId"`
	Name                  string   `json:"name"`
	Color                 string   `json:"color"`
	Group                 Group    `json:"group"`
	IsCoreEvent           bool     `json:"isCoreEvent"`
	ValidThrough          int64    `json:"validThrough"`
	Dirty                 bool     `json:"dirty"`
	DailyMergeFirst       int      `json:"dailyMergeFirst"`
	DailyRollupFirst      int      `json:"dailyRollupFirst"`
	TrackTypeName         string   `json:"trackTypeName"`
	TrackTypeRules        []string `json:"trackTypeRules"`
	EventPropertyNameList []string `json:"eventPropertyNameList"`
}

type Step struct {
	Id                          string         `json:"id"`
	GuideId                     string         `json:"guideId"`
	TemplateId                  string         `json:"templateId"`
	Type                        string         `json:"type"`
	ElementPathRule             string         `json:"elementPathRule"`
	TriggerElementPathRule      string         `json:"triggerElementPathRule"`
	ConfirmationElementPathRule string         `json:"confirmationElementPathRule"`
	ContentType                 string         `json:"contentType"`
	BuildingBlocksUrl           string         `json:"buildingBlocksUrl"`
	DomUrl                      string         `json:"domUrl"`
	DomJsonpUrl                 string         `json:"domJsonpUrl"`
	Rank                        int            `json:"rank"`
	AdvanceMethod               string         `json:"advanceMethod"`
	Attributes                  StepAttributes `json:"attributes"`
	LastUpdatedAt               int64          `json:"lastUpdatedAt"`
	ResetAt                     int            `json:"resetAt"`
	LaunchUrl                   string         `json:"launchUrl,omitempty"`
	PageId                      string         `json:"pageId,omitempty"`
	RegexUrlRule                string         `json:"regexUrlRule,omitempty"`
	PollIds                     []string       `json:"pollIds,omitempty"`
}

type StepAttributes struct {
	Css                  string         `json:"css,omitempty"`
	ThemeId              string         `json:"themeId"`
	IsAutoFocus          bool           `json:"isAutoFocus"`
	BlockOutUI           BlockOutUI     `json:"blockOutUI"`
	ElementSelectionType string         `json:"elementSelectionType,omitempty"`
	LayoutDir            string         `json:"layoutDir,omitempty"`
	AdvanceActions       AdvanceActions `json:"advanceActions,omitempty"`
}

type BlockOutUI struct {
	AdditionalElements string  `json:"additionalElements"`
	Enabled            bool    `json:"enabled"`
	Padding            Padding `json:"padding"`
}

type AdvanceActions struct {
	ElementClick bool `json:"elementClick"`
	ElementHover bool `json:"elementHover"`
}

type Padding struct {
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
	Right  int `json:"right"`
	Top    int `json:"top"`
}

type Guide struct {
	CreatedByUser     User            `json:"createdByUser"`
	CreatedAt         int64           `json:"createdAt"`
	LastUpdatedByUser User            `json:"lastUpdatedByUser"`
	LastUpdatedAt     int64           `json:"lastUpdatedAt"`
	Kind              string          `json:"kind"`
	RootVersionId     string          `json:"rootVersionId"`
	StableVersionId   string          `json:"stableVersionId"`
	AppId             int             `json:"appId"`
	AppIds            []int           `json:"appIds"`
	Id                string          `json:"id"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	State             string          `json:"state"`
	EmailState        string          `json:"emailState"`
	LaunchMethod      string          `json:"launchMethod"`
	IsMultiStep       bool            `json:"isMultiStep"`
	IsTraining        bool            `json:"isTraining"`
	Steps             []Step          `json:"steps"`
	Attributes        GuideAttributes `json:"attributes"`
	Audience          []struct {
		Source struct {
			Visitors any `json:"visitors"`
		} `json:"source,omitempty"`
		Eval struct {
			AccountId string `json:"accountId"`
		} `json:"eval,omitempty"`
		Unwind struct {
			Field     string `json:"field"`
			KeepEmpty bool   `json:"keepEmpty"`
		} `json:"unwind,omitempty"`
		Segment struct {
			Id string `json:"id"`
		} `json:"segment,omitempty"`
		Select struct {
			VisitorId string `json:"visitorId"`
		} `json:"select,omitempty"`
		Identified string `json:"identified,omitempty"`
		Filter     string `json:"filter,omitempty"`
	} `json:"audience"`
	AudienceUiHint struct {
		Filters []struct {
			Kind        string `json:"kind"`
			SegmentId   string `json:"segmentId,omitempty"`
			Key         string `json:"key,omitempty"`
			Group       string `json:"group,omitempty"`
			Field       string `json:"field,omitempty"`
			Id          string `json:"id,omitempty"`
			Type        string `json:"type,omitempty"`
			Name        string `json:"name,omitempty"`
			Schema      string `json:"schema,omitempty"`
			ElementType string `json:"elementType,omitempty"`
			Operator    string `json:"operator,omitempty"`
			Value       string `json:"value,omitempty"`
			First       any    `json:"first"`
			Last        any    `json:"last"`
			Valid       bool   `json:"$valid,omitempty"`
		} `json:"filters"`
	} `json:"audienceUiHint"`
	AuthoredLanguage               string   `json:"authoredLanguage"`
	Recurrence                     int      `json:"recurrence"`
	RecurrenceEligibilityWindow    int      `json:"recurrenceEligibilityWindow"`
	ResetAt                        int      `json:"resetAt"`
	PublishedAt                    int64    `json:"publishedAt"`
	PublishedEver                  bool     `json:"publishedEver"`
	CurrentFirstEligibleToBeSeenAt int      `json:"currentFirstEligibleToBeSeenAt"`
	ExpiresAfter                   int64    `json:"expiresAfter,omitempty"`
	IsTopLevel                     bool     `json:"isTopLevel"`
	IsModule                       bool     `json:"isModule"`
	EditorType                     string   `json:"editorType"`
	DependentMetadata              []string `json:"dependentMetadata"`
	ShowsAfter                     int64    `json:"showsAfter,omitempty"`
	Polls                          []struct {
		Id               string `json:"id"`
		Question         string `json:"question"`
		NumericResponses []int  `json:"numericResponses"`
		IdResponses      struct {
			Field1 string `json:"0"`
			Field2 string `json:"1"`
		} `json:"idResponses"`
		Attributes struct {
			Type string `json:"type"`
		} `json:"attributes"`
		ResetAt int `json:"resetAt"`
	} `json:"polls,omitempty"`
	TranslationStates struct {
		EnGB struct {
			State                 string    `json:"state"`
			Message               string    `json:"message"`
			TranslationObjectName string    `json:"translationObjectName"`
			TranslationAppliedAt  time.Time `json:"translationAppliedAt"`
			StepTranslations      struct {
				CbtRNWIVmsEmufn0LVvjiDVp4A struct {
					TranslatedStringIds   []any    `json:"translatedStringIds"`
					UntranslatedStringIds []string `json:"untranslatedStringIds"`
					TranslationHash       string   `json:"translationHash"`
					DomHash               string   `json:"domHash"`
					DomHash256            string   `json:"domHash256"`
					DomJsonpHash          string   `json:"domJsonpHash"`
					DomJsonpHash256       string   `json:"domJsonpHash256"`
				} `json:"CbtR_NWIVmsEmufn0LVvjiDVp4A"`
				DnDsXMfqjMvZz30B9HWw3RO3BM struct {
					TranslatedStringIds   []any    `json:"translatedStringIds"`
					UntranslatedStringIds []string `json:"untranslatedStringIds"`
					TranslationHash       string   `json:"translationHash"`
					DomHash               string   `json:"domHash"`
					DomHash256            string   `json:"domHash256"`
					DomJsonpHash          string   `json:"domJsonpHash"`
					DomJsonpHash256       string   `json:"domJsonpHash256"`
				} `json:"dnDsXMfqjMvZz30b9hWw3_rO3BM"`
			} `json:"stepTranslations"`
		} `json:"en-GB"`
	} `json:"translationStates,omitempty"`
}

type GuideAttributes struct {
	Dates struct {
		EnGB string `json:"en-GB"`
		EnUS string `json:"en-US"`
		Es   string `json:"es"`
		Nl   string `json:"nl"`
	} `json:"dates"`
	Device struct {
		Type string `json:"type"`
	} `json:"device"`
	Priority             int    `json:"priority"`
	SharedServiceVersion string `json:"sharedServiceVersion"`
	Type                 string `json:"type"`
	Activation           struct {
		Event          []string `json:"event"`
		InheritStepOne bool     `json:"inheritStepOne"`
		Selector       string   `json:"selector"`
	} `json:"activation,omitempty"`
	Capping struct {
		MaxImpressions        int `json:"maxImpressions"`
		MaxSessionImpressions int `json:"maxSessionImpressions"`
	} `json:"capping,omitempty"`
	Dom struct {
		IsOnlyShowOnce              bool `json:"isOnlyShowOnce"`
		ShowGuideOnlyOnElementClick bool `json:"showGuideOnlyOnElementClick"`
	} `json:"dom,omitempty"`
	ElementSelectionType string `json:"elementSelectionType,omitempty"`
}
type Report struct {
	CreatedByUser     User   `json:"createdByUser"`
	CreatedAt         int64  `json:"createdAt"`
	LastUpdatedByUser User   `json:"lastUpdatedByUser"`
	LastUpdatedAt     int64  `json:"lastUpdatedAt"`
	Kind              string `json:"kind"`
	RootVersionId     string `json:"rootVersionId"`
	StableVersionId   string `json:"stableVersionId"`
	OwnedByUser       User   `json:"ownedByUser"`
	Share             string `json:"share"`
	Target            string `json:"target"`
	Scope             string `json:"scope"`
	Level             string `json:"level"`
	Id                string `json:"id"`
	Type              string `json:"type"`
	Name              string `json:"name"`
	Shared            bool   `json:"shared"`
	Definition        struct {
		Config struct {
			MaxDuration int `json:"maxDuration,omitempty"`
			MaxInterval int `json:"maxInterval,omitempty"`
			Items       []struct {
				PageId string `json:"pageId"`
			} `json:"items,omitempty"`
			AppId           any    `json:"appId"`
			SegmentId       string `json:"segmentId"`
			SelectedAccount struct {
				Id        string `json:"id"`
				Text      string `json:"text"`
				ClassName string `json:"className"`
			} `json:"selectedAccount,omitempty"`
			Columns []struct {
				ColumnLabel  string   `json:"columnLabel"`
				ColumnName   string   `json:"columnName"`
				FieldName    string   `json:"fieldName"`
				FunctionName string   `json:"functionName"`
				Immutable    bool     `json:"immutable"`
				Kind         string   `json:"kind"`
				PrimaryKeys  []string `json:"primaryKeys"`
				Schema       string   `json:"schema"`
				Disabled     bool     `json:"disabled"`
			} `json:"columns,omitempty"`
			Filters        []any  `json:"filters,omitempty"`
			Segments       []any  `json:"segments,omitempty"`
			Predecessors   bool   `json:"predecessors,omitempty"`
			OmitPages      bool   `json:"omitPages,omitempty"`
			Features       bool   `json:"features,omitempty"`
			TrackEvents    bool   `json:"trackEvents,omitempty"`
			CollapseDups   bool   `json:"collapseDups,omitempty"`
			MaxLength      int    `json:"maxLength,omitempty"`
			PageId         string `json:"pageId,omitempty"`
			UniqueVisitors bool   `json:"uniqueVisitors,omitempty"`
		} `json:"config,omitempty"`
		TimeSeries struct {
			Type  string `json:"type"`
			Start any    `json:"start"`
			End   any    `json:"end"`
		} `json:"timeSeries,omitempty"`
		Type         string  `json:"type,omitempty"`
		History      []int64 `json:"history,omitempty"`
		Version      int     `json:"version,omitempty"`
		Generator    string  `json:"generator,omitempty"`
		Kind         string  `json:"kind,omitempty"`
		TargetGroups []struct {
			Id struct {
				Value string `json:"value"`
				Valid bool   `json:"valid"`
			} `json:"id"`
			Measure struct {
				Value struct {
					Name                string `json:"name"`
					Label               string `json:"label"`
					DisabledTargetTypes []any  `json:"disabledTargetTypes"`
					Disabled            bool   `json:"disabled"`
				} `json:"value"`
				Valid bool `json:"valid"`
			} `json:"measure"`
			Targets []struct {
				Id      int    `json:"id"`
				Type    string `json:"type"`
				Filters []any  `json:"filters"`
			} `json:"targets"`
		} `json:"targetGroups,omitempty"`
		Globals struct {
			DateConfigs []struct {
				Value struct {
					DateRange struct {
						Type string `json:"type"`
					} `json:"dateRange"`
					Period   string `json:"period"`
					Interval string `json:"interval"`
				} `json:"value"`
				IsPrimary bool `json:"isPrimary"`
				Valid     bool `json:"valid"`
			} `json:"dateConfigs"`
			Segments []struct {
				Value struct {
					Order      int    `json:"order"`
					Default    bool   `json:"default"`
					Name       string `json:"name"`
					Id         string `json:"id"`
					Shared     bool   `json:"shared"`
					Definition string `json:"definition"`
					Disabled   any    `json:"disabled"`
				} `json:"value"`
				Valid bool `json:"valid"`
			} `json:"segments"`
			Groupings []struct {
				Configuration struct {
					Type             string `json:"Type"`
					DisplayName      string `json:"DisplayName"`
					ElementType      string `json:"ElementType"`
					ElementFormat    string `json:"ElementFormat"`
					Dirty            bool   `json:"Dirty"`
					Cardinality      int    `json:"cardinality"`
					IsHidden         bool   `json:"isHidden"`
					IsDeleted        bool   `json:"isDeleted"`
					IsCalculated     bool   `json:"isCalculated"`
					IsPerApp         bool   `json:"isPerApp"`
					NeverIndex       bool   `json:"neverIndex"`
					Group            string `json:"group"`
					Field            string `json:"field"`
					Kind             string `json:"kind"`
					GroupDisplayName string `json:"groupDisplayName"`
					IsPromoted       bool   `json:"isPromoted"`
				} `json:"configuration"`
				Type  string `json:"type"`
				Valid bool   `json:"valid"`
			} `json:"groupings"`
			Filters []struct {
				Field            string   `json:"field"`
				Kind             string   `json:"kind"`
				FilterType       string   `json:"filterType"`
				Operator         string   `json:"operator"`
				ComparisonValues []string `json:"comparisonValues"`
				Id               string   `json:"id"`
			} `json:"filters"`
			Formulas []any `json:"formulas"`
		} `json:"globals,omitempty"`
		ChartSelections struct {
			Bde2CEc93483898F4F9F41Eacc0Ff struct {
				Type string `json:"type"`
			} `json:"972bde2c-ec93-4838-98f4-f9f41eacc0ff"`
			Globals struct {
				MatchScales bool `json:"matchScales"`
			} `json:"globals"`
		} `json:"chartSelections,omitempty"`
		RowSelections struct {
			Bde2CEc93483898F4F9F41Eacc0Ff323232TestVisitor1      bool `json:"972bde2c-ec93-4838-98f4-f9f41eacc0ff--323232-testVisitor1"`
			Bde2CEc93483898F4F9F41Eacc0Ff323232OhHaiImAVisitor   bool `json:"972bde2c-ec93-4838-98f4-f9f41eacc0ff--323232-OhHai_ImAVisitor"`
			Bde2CEc93483898F4F9F41Eacc0Ff323232PENDOTAaVUTi1CQOd bool `json:"972bde2c-ec93-4838-98f4-f9f41eacc0ff--323232-_PENDO_T_aaVUTi1CQOd"`
			Bde2CEc93483898F4F9F41Eacc0Ff323232NewSessionVisitor bool `json:"972bde2c-ec93-4838-98f4-f9f41eacc0ff--323232-newSessionVisitor"`
			Bde2CEc93483898F4F9F41Eacc0Ff323232TestVisitor2      bool `json:"972bde2c-ec93-4838-98f4-f9f41eacc0ff--323232-testVisitor2"`
		} `json:"rowSelections,omitempty"`
		Minimum int `json:"minimum,omitempty"`
	} `json:"definition"`
	Aggregation struct {
		Pipeline []struct {
			Source struct {
				SingleEvents struct {
					AppId       any      `json:"appId"`
					Blacklist   string   `json:"blacklist"`
					ReverseTime bool     `json:"reverseTime,omitempty"`
					EventClass  []string `json:"eventClass,omitempty"`
				} `json:"singleEvents,omitempty"`
				TimeSeries struct {
					Period string `json:"period"`
					First  any    `json:"first"`
					Count  int    `json:"count,omitempty"`
					Last   string `json:"last,omitempty"`
				} `json:"timeSeries"`
				Events struct {
					AppId     any    `json:"appId"`
					Blacklist string `json:"blacklist"`
				} `json:"events,omitempty"`
			} `json:"source,omitempty"`
			Group struct {
				Group  []string `json:"group"`
				Fields []struct {
					Funnel struct {
						Funnel struct {
							Blacklist        string `json:"blacklist"`
							IncludeAnonymous any    `json:"includeAnonymous"`
							MaxDuration      int    `json:"maxDuration"`
							MaxInterval      int    `json:"maxInterval"`
							Items            []struct {
								PageId string `json:"pageId"`
							} `json:"items"`
							AppId     []int `json:"appId"`
							SegmentId any   `json:"segmentId"`
						} `json:"funnel"`
					} `json:"funnel,omitempty"`
					Path struct {
						Path struct {
							PageId               string `json:"pageId"`
							CollapseDups         bool   `json:"collapseDups"`
							MaxInterval          int    `json:"maxInterval"`
							MaxLength            int    `json:"maxLength"`
							Predecessors         bool   `json:"predecessors"`
							FollowAcrossSessions bool   `json:"followAcrossSessions"`
							OmitPages            bool   `json:"omitPages"`
							Features             bool   `json:"features"`
							TrackEvents          bool   `json:"trackEvents"`
						} `json:"path"`
					} `json:"path,omitempty"`
				} `json:"fields,omitempty"`
			} `json:"group,omitempty"`
			Unwind struct {
				Field string `json:"field"`
			} `json:"unwind,omitempty"`
			Select struct {
				Times                string `json:"times,omitempty"`
				Start                string `json:"start,omitempty"`
				Steps                string `json:"steps,omitempty"`
				VisitorId            string `json:"visitorId,omitempty"`
				VisitorAgentName     string `json:"visitor_agent_name,omitempty"`
				AccountId            string `json:"accountId,omitempty"`
				AccountAutoLastvisit string `json:"account_auto_lastvisit,omitempty"`
			} `json:"select,omitempty"`
			Fork [][]struct {
				Group struct {
					Group  []string `json:"group"`
					Fields []struct {
						Count struct {
							Count *string `json:"count"`
						} `json:"count"`
					} `json:"fields"`
				} `json:"group,omitempty"`
				Sort       []string `json:"sort,omitempty"`
				Accumulate struct {
					Count string `json:"count"`
				} `json:"accumulate,omitempty"`
				Compute struct {
					Durations struct {
						Diff string `json:"diff"`
					} `json:"durations"`
				} `json:"compute,omitempty"`
				Reduce struct {
					AverageTimes struct {
						ListAverage string `json:"listAverage"`
					} `json:"averageTimes,omitempty"`
					Counts struct {
						List string `json:"list"`
					} `json:"counts,omitempty"`
					FurthestStep struct {
						Max string `json:"max"`
					} `json:"furthestStep,omitempty"`
					AverageTimeToCompletion struct {
						Sum string `json:"sum"`
					} `json:"averageTimeToCompletion,omitempty"`
				} `json:"reduce,omitempty"`
				Eval struct {
					PercentCompleted string `json:"percentCompleted,omitempty"`
					Items            string `json:"items,omitempty"`
				} `json:"eval,omitempty"`
				Filter string `json:"filter,omitempty"`
				Unwind struct {
					Field string `json:"field"`
				} `json:"unwind,omitempty"`
				Limit   int `json:"limit,omitempty"`
				Treeify struct {
					Threshold float64 `json:"threshold"`
					KeySort   bool    `json:"keySort"`
				} `json:"treeify,omitempty"`
			} `json:"fork,omitempty"`
			BulkExpand struct {
				Visitor struct {
					Visitor string `json:"visitor"`
				} `json:"visitor,omitempty"`
				Account struct {
					Account string `json:"account"`
				} `json:"account,omitempty"`
			} `json:"bulkExpand,omitempty"`
			Eval struct {
				VisitorAgentName     string `json:"visitor_agent_name,omitempty"`
				AccountAutoLastvisit string `json:"account_auto_lastvisit,omitempty"`
			} `json:"eval,omitempty"`
			Segment struct {
				Id string `json:"id"`
			} `json:"segment,omitempty"`
		} `json:"pipeline,omitempty"`
		Fields []struct {
			Type  string `json:"type"`
			Title string `json:"title"`
			Field string `json:"field"`
		} `json:"fields,omitempty"`
	} `json:"aggregation"`
	LastRunAt         int64  `json:"lastRunAt"`
	LastSuccessRunAt  int64  `json:"lastSuccessRunAt,omitempty"`
	LastSuccessRunObj string `json:"lastSuccessRunObj,omitempty"`
	AggregationList   []struct {
		Pipeline []struct {
			Source struct {
				Events struct {
					AppId     []int  `json:"appId"`
					Blacklist string `json:"blacklist"`
				} `json:"events"`
				TimeSeries struct {
					Period string `json:"period"`
					First  string `json:"first"`
					Count  int    `json:"count"`
				} `json:"timeSeries"`
			} `json:"source,omitempty"`
			BulkExpand struct {
				Visitor struct {
					Visitor string `json:"visitor"`
				} `json:"visitor"`
			} `json:"bulkExpand,omitempty"`
			Useragent struct {
				U string `json:"u"`
			} `json:"useragent,omitempty"`
			Eval struct {
				GroupValue0         string `json:"groupValue0,omitempty"`
				NameDeviceAutomatic string `json:"name_device_automatic,omitempty"`
				Month               string `json:"month,omitempty"`
			} `json:"eval,omitempty"`
			Filter string `json:"filter,omitempty"`
			Group  struct {
				Group  []string `json:"group"`
				Fields []struct {
					VisitorCount struct {
						Count string `json:"count"`
					} `json:"visitorCount"`
				} `json:"fields"`
			} `json:"group,omitempty"`
			Limit  int `json:"limit,omitempty"`
			Select struct {
				Month       string `json:"month"`
				AppId       string `json:"appId"`
				YAxis       string `json:"yAxis"`
				GroupValue0 string `json:"groupValue0"`
			} `json:"select,omitempty"`
		} `json:"pipeline"`
	} `json:"aggregationList,omitempty"`
}
