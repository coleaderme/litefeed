package workers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type TimelineMedia struct {
	Caption string
	Image   string
	Video   string
}
type PageData struct {
	Media     []TimelineMedia
	EndCursor string
	Username  string
}
type ProfileJson struct {
	Data struct {
		User struct {
			FriendshipStatus         interface{}   `json:"friendship_status"`
			Fullname                 string        `json:"full_name"`
			Gating                   interface{}   `json:"gating"`
			IsCheckpointMemorialized bool          `json:"is_checkpoint_memorialized"`
			IsPrivate                bool          `json:"is_private"`
			HasStoryArchive          interface{}   `json:"has_story_archive"`
			Username                 string        `json:"username"`
			SupervisionInfo          interface{}   `json:"supervision_info"`
			IsRegulatedC18           bool          `json:"is_regulated_c18"`
			RegulatedNewsInLocations []interface{} `json:"regulated_news_in_locations"`
			TextPostAppBadgeLabel    interface{}   `json:"text_post_app_badge_label"`
			ShowTextPostAppBadge     interface{}   `json:"show_text_post_app_badge"`
			TextPostNewPostCount     interface{}   `json:"text_post_new_post_count"`
			Pk                       string        `json:"pk"`
			LiveBroadcastVisibility  interface{}   `json:"live_broadcast_visibility"`
			LiveBroadcastID          interface{}   `json:"live_broadcast_id"`
			ProfilePicURL            string        `json:"profile_pic_url"`
			HdProfilePicURLInfo      struct {
				URL string `json:"url"`
			} `json:"hd_profile_pic_url_info"`
			IsUnpublished                  bool          `json:"is_unpublished"`
			LatestReelMedia                interface{}   `json:"latest_reel_media"`
			HasProfilePic                  interface{}   `json:"has_profile_pic"`
			AccountType                    int           `json:"account_type"`
			FollowerCount                  int           `json:"follower_count"`
			IsVerified                     bool          `json:"is_verified"`
			MutualFollowersCount           interface{}   `json:"mutual_followers_count"`
			ProfileContextLinksWithUserIds interface{}   `json:"profile_context_links_with_user_ids"`
			AddressStreet                  interface{}   `json:"address_street"`
			CityName                       interface{}   `json:"city_name"`
			IsBusiness                     bool          `json:"is_business"`
			Zip                            interface{}   `json:"zip"`
			BiographyWithEntities          interface{}   `json:"biography_with_entities"`
			Category                       interface{}   `json:"category"`
			ShouldShowCategory             interface{}   `json:"should_show_category"`
			AccountBadges                  []interface{} `json:"account_badges"`
			AiAgentType                    interface{}   `json:"ai_agent_type"`
			BioLinks                       []interface{} `json:"bio_links"`
			FbProfileBioLinkWeb            interface{}   `json:"fb_profile_bio_link_web"`
			ExternalLynxURL                interface{}   `json:"external_lynx_url"`
			ExternalURL                    string        `json:"external_url"`
			Pronouns                       []interface{} `json:"pronouns"`
			Biography                      string        `json:"biography"`
			TransparencyLabel              interface{}   `json:"transparency_label"`
			TransparencyProduct            interface{}   `json:"transparency_product"`
			HasChaining                    interface{}   `json:"has_chaining"`
			RemoveMessageEntrypoint        interface{}   `json:"remove_message_entrypoint"`
			FbidV2                         string        `json:"fbid_v2"`
			InteropMessagingUserFbid       string        `json:"interop_messaging_user_fbid"`
			ShowAccountTransparencyDetails bool          `json:"show_account_transparency_details"`
			IsEmbedsDisabled               bool          `json:"is_embeds_disabled"`
			IsProfessionalAccount          interface{}   `json:"is_professional_account"`
			FollowingCount                 int           `json:"following_count"`
			MediaCount                     int           `json:"media_count"`
			TotalClipsCount                interface{}   `json:"total_clips_count"`
			LatestBestiesReelMedia         interface{}   `json:"latest_besties_reel_media"`
			ReelMediaSeenTimestamp         interface{}   `json:"reel_media_seen_timestamp"`
			ID                             string        `json:"id"`
		} `json:"user"`
		Viewer struct {
			User interface{} `json:"user"`
		} `json:"viewer"`
	} `json:"data"`
	Extensions struct {
		IsFinal bool `json:"is_final"`
	} `json:"extensions"`
	Status string `json:"status"`
}
type TimelineFirstJson struct {
	Data struct {
		XdtAPIV1FeedUserTimelineGraphqlConnection struct {
			Edges []struct {
				Node struct {
					Code                       string      `json:"code"`
					Pk                         string      `json:"pk"`
					ID                         string      `json:"id"`
					AdID                       interface{} `json:"ad_id"`
					BoostedStatus              interface{} `json:"boosted_status"`
					BoostUnavailableIdentifier interface{} `json:"boost_unavailable_identifier"`
					BoostUnavailableReason     interface{} `json:"boost_unavailable_reason"`
					Caption                    struct {
						HasTranslation interface{} `json:"has_translation"`
						CreatedAt      int         `json:"created_at"`
						Pk             string      `json:"pk"`
						Text           string      `json:"text"`
					} `json:"caption"`
					CaptionIsEdited         bool        `json:"caption_is_edited"`
					FeedDemotionControl     interface{} `json:"feed_demotion_control"`
					FeedRecsDemotionControl interface{} `json:"feed_recs_demotion_control"`
					TakenAt                 int         `json:"taken_at"`
					InventorySource         interface{} `json:"inventory_source"`
					VideoVersions           []struct {
						Width  int    `json:"width"`
						Height int    `json:"height"`
						URL    string `json:"url"`
						Type   int    `json:"type"`
					} `json:"video_versions"`
					IsDashEligible    int    `json:"is_dash_eligible"`
					NumberOfQualities int    `json:"number_of_qualities"`
					VideoDashManifest string `json:"video_dash_manifest"`
					ImageVersions2    struct {
						Candidates []struct {
							URL    string `json:"url"`
							Height int    `json:"height"`
							Width  int    `json:"width"`
						} `json:"candidates"`
					} `json:"image_versions2"`
					IsPaidPartnership    bool        `json:"is_paid_partnership"`
					SponsorTags          interface{} `json:"sponsor_tags"`
					AffiliateInfo        interface{} `json:"affiliate_info"`
					OriginalHeight       int         `json:"original_height"`
					OriginalWidth        int         `json:"original_width"`
					OrganicTrackingToken string      `json:"organic_tracking_token"`
					Link                 interface{} `json:"link"`
					StoryCta             interface{} `json:"story_cta"`
					User                 struct {
						Pk               string      `json:"pk"`
						Username         string      `json:"username"`
						ProfilePicURL    string      `json:"profile_pic_url"`
						IsPrivate        bool        `json:"is_private"`
						IsEmbedsDisabled interface{} `json:"is_embeds_disabled"`
						IsUnpublished    bool        `json:"is_unpublished"`
						IsVerified       bool        `json:"is_verified"`
						FriendshipStatus struct {
							Following      bool `json:"following"`
							IsBestie       bool `json:"is_bestie"`
							IsFeedFavorite bool `json:"is_feed_favorite"`
							IsRestricted   bool `json:"is_restricted"`
						} `json:"friendship_status"`
						LatestBestiesReelMedia  int         `json:"latest_besties_reel_media"`
						LatestReelMedia         int         `json:"latest_reel_media"`
						LiveBroadcastVisibility interface{} `json:"live_broadcast_visibility"`
						LiveBroadcastID         interface{} `json:"live_broadcast_id"`
						Seen                    interface{} `json:"seen"`
						SupervisionInfo         interface{} `json:"supervision_info"`
						ID                      string      `json:"id"`
						HdProfilePicURLInfo     struct {
							URL string `json:"url"`
						} `json:"hd_profile_pic_url_info"`
						FullName string `json:"full_name"`
						Typename string `json:"__typename"`
					} `json:"user"`
					Group interface{} `json:"group"`
					Owner struct {
						Pk               string `json:"pk"`
						ProfilePicURL    string `json:"profile_pic_url"`
						Username         string `json:"username"`
						FriendshipStatus struct {
							IsFeedFavorite bool `json:"is_feed_favorite"`
							Following      bool `json:"following"`
							IsRestricted   bool `json:"is_restricted"`
							IsBestie       bool `json:"is_bestie"`
						} `json:"friendship_status"`
						IsEmbedsDisabled               interface{} `json:"is_embeds_disabled"`
						IsUnpublished                  bool        `json:"is_unpublished"`
						IsVerified                     bool        `json:"is_verified"`
						ShowAccountTransparencyDetails bool        `json:"show_account_transparency_details"`
						SupervisionInfo                interface{} `json:"supervision_info"`
						TransparencyProduct            interface{} `json:"transparency_product"`
						TransparencyProductEnabled     bool        `json:"transparency_product_enabled"`
						TransparencyLabel              interface{} `json:"transparency_label"`
						ID                             string      `json:"id"`
						Typename                       string      `json:"__typename"`
						IsPrivate                      bool        `json:"is_private"`
					} `json:"owner"`
					CoauthorProducers           []interface{} `json:"coauthor_producers"`
					InvitedCoauthorProducers    []interface{} `json:"invited_coauthor_producers"`
					FollowHashtagInfo           interface{}   `json:"follow_hashtag_info"`
					Title                       interface{}   `json:"title"`
					CommentCount                int           `json:"comment_count"`
					CommentsDisabled            interface{}   `json:"comments_disabled"`
					CommentingDisabledForViewer interface{}   `json:"commenting_disabled_for_viewer"`
					LikeAndViewCountsDisabled   bool          `json:"like_and_view_counts_disabled"`
					HasLiked                    bool          `json:"has_liked"`
					TopLikers                   []string      `json:"top_likers"`
					FacepileTopLikers           []struct {
						ProfilePicURL string `json:"profile_pic_url"`
						Pk            string `json:"pk"`
						Username      string `json:"username"`
						ID            string `json:"id"`
					} `json:"facepile_top_likers"`
					LikeCount             int         `json:"like_count"`
					Preview               interface{} `json:"preview"`
					CanSeeInsightsAsBrand bool        `json:"can_see_insights_as_brand"`
					SocialContext         []struct {
						SocialContextType          string `json:"social_context_type"`
						SocialContextUsersCount    int    `json:"social_context_users_count"`
						SocialContextFacepileUsers []struct {
							ProfilePicURL string `json:"profile_pic_url"`
							PkID          string `json:"pk_id"`
							Username      string `json:"username"`
							ID            string `json:"id"`
							UserID        string `json:"userID"`
						} `json:"social_context_facepile_users"`
						Typename string `json:"__typename"`
					} `json:"social_context"`
					ViewCount              interface{} `json:"view_count"`
					CanReshare             interface{} `json:"can_reshare"`
					CanViewerReshare       bool        `json:"can_viewer_reshare"`
					IgMediaSharingDisabled bool        `json:"ig_media_sharing_disabled"`
					PhotoOfYou             interface{} `json:"photo_of_you"`
					ProductType            string      `json:"product_type"`
					MediaType              int         `json:"media_type"`
					Usertags               interface{} `json:"usertags"`
					MediaOverlayInfo       interface{} `json:"media_overlay_info"`
					CarouselParentID       interface{} `json:"carousel_parent_id"`
					CarouselMediaCount     interface{} `json:"carousel_media_count"`
					CarouselMedia          interface{} `json:"carousel_media"`
					Location               interface{} `json:"location"`
					HasAudio               bool        `json:"has_audio"`
					ClipsMetadata          struct {
						AudioType        string `json:"audio_type"`
						AchievementsInfo struct {
							ShowAchievements bool `json:"show_achievements"`
						} `json:"achievements_info"`
						MusicInfo         interface{} `json:"music_info"`
						OriginalSoundInfo struct {
							OriginalAudioTitle string `json:"original_audio_title"`
							ShouldMuteAudio    bool   `json:"should_mute_audio"`
							AudioAssetID       string `json:"audio_asset_id"`
							ConsumptionInfo    struct {
								IsTrendingInClips         bool        `json:"is_trending_in_clips"`
								ShouldMuteAudioReason     string      `json:"should_mute_audio_reason"`
								ShouldMuteAudioReasonType interface{} `json:"should_mute_audio_reason_type"`
							} `json:"consumption_info"`
							IgArtist struct {
								Username string `json:"username"`
								ID       string `json:"id"`
							} `json:"ig_artist"`
							IsExplicit bool `json:"is_explicit"`
						} `json:"original_sound_info"`
					} `json:"clips_metadata"`
					ClipsAttributionInfo interface{} `json:"clips_attribution_info"`
					AccessibilityCaption interface{} `json:"accessibility_caption"`
					Audience             interface{} `json:"audience"`
					DisplayURI           interface{} `json:"display_uri"`
					MediaCroppingInfo    struct {
						SquareCrop struct {
							CropBottom float64 `json:"crop_bottom"`
							CropLeft   float64 `json:"crop_left"`
							CropRight  float64 `json:"crop_right"`
							CropTop    float64 `json:"crop_top"`
						} `json:"square_crop"`
					} `json:"media_cropping_info"`
					Thumbnails                      interface{} `json:"thumbnails"`
					TimelinePinnedUserIds           []string    `json:"timeline_pinned_user_ids"`
					UpcomingEvent                   interface{} `json:"upcoming_event"`
					Explore                         interface{} `json:"explore"`
					MainFeedCarouselStartingMediaID interface{} `json:"main_feed_carousel_starting_media_id"`
					IsSeen                          interface{} `json:"is_seen"`
					OpenCarouselSubmissionState     interface{} `json:"open_carousel_submission_state"`
					PreviousSubmitter               interface{} `json:"previous_submitter"`
					AllPreviousSubmitters           interface{} `json:"all_previous_submitters"`
					Headline                        interface{} `json:"headline"`
					Comments                        interface{} `json:"comments"`
					SavedCollectionIds              interface{} `json:"saved_collection_ids"`
					HasViewerSaved                  interface{} `json:"has_viewer_saved"`
					SharingFrictionInfo             struct {
						ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
						BloksAppURL               interface{} `json:"bloks_app_url"`
					} `json:"sharing_friction_info"`
					MediaLevelCommentControls interface{} `json:"media_level_comment_controls"`
					Typename                  string      `json:"__typename"`
				} `json:"node"`
				Cursor string `json:"cursor"`
			} `json:"edges"`
			PageInfo struct {
				EndCursor       string      `json:"end_cursor"`
				HasNextPage     bool        `json:"has_next_page"`
				HasPreviousPage bool        `json:"has_previous_page"`
				StartCursor     interface{} `json:"start_cursor"`
			} `json:"page_info"`
		} `json:"xdt_api__v1__feed__user_timeline_graphql_connection"`
		XdtViewer struct {
			User struct {
				ID string `json:"id"`
			} `json:"user"`
		} `json:"xdt_viewer"`
	} `json:"data"`
	Extensions struct {
		IsFinal bool `json:"is_final"`
	} `json:"extensions"`
	Status string `json:"status"`
}
type AutoGenerated struct {
	Data struct {
		XdtAPIV1FbsearchTopsearchConnection struct {
			Users []struct {
				User struct {
					Username      string `json:"username"`
					Fullname      string `json:"full_name"`
					ID            string `json:"id"`
					ProfilePicURL string `json:"profile_pic_url"`
				} `json:"user"`
			} `json:"users"`
		} `json:"xdt_api__v1__fbsearch__topsearch_connection"`
	} `json:"data"`
}
type UserProfile struct {
	Fullname   string `json:"full_name"`
	Username   string `json:"username"`
	Biography  string `json:"biography"`
	URL        string `json:"url"`
	IsVerified bool   `json:"is_verified"`
}
type User struct {
	Fullname      string `json:"full_name"`
	Username      string `json:"username"`
	ID            string `json:"id"`
	ProfilePicURL string `json:"profile_pic_url"`
}

func Profile(id string, client *http.Client) (UserProfile, error) {
	// Search and Profile doesnt send same data :(

	// ADD DATA STRING HERE
	var data = strings.NewReader(``)
	req, err := http.NewRequest("POST", "https://www.instagram.com/graphql/query", data)
	if err != nil {
		log.Fatal(err)
	}
	// ADD HEADERS HERE

	// Request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return UserProfile{}, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// JSON parse
	var result ProfileJson
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	// Extracting the values
	var username string = result.Data.User.Username
	var fullName string = result.Data.User.Fullname
	var biography string = result.Data.User.Biography
	var hdProfilePicURL string = result.Data.User.HdProfilePicURLInfo.URL
	var is_verified bool = result.Data.User.IsVerified
	user := UserProfile{
		Fullname:   fullName,
		Username:   username,
		Biography:  biography,
		IsVerified: is_verified,
		URL:        "/proxy?url=" + url.QueryEscape(hdProfilePicURL),
	}
	return user, nil
}

func Search(query string, client *http.Client) ([]User, error) {
	// Search and Profile doesnt send same data :(

	// DATA STRING HERE | NOT THE SAME
	data := strings.NewReader("")
	req, err := http.NewRequest("POST", "https://www.instagram.com/graphql/query", data)
	if err != nil {
		log.Fatal(err)
	}
	// ADD HEADERS HERE | HEADERS ARE SAME

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// JSON parse
	var autoGen AutoGenerated
	if err := json.Unmarshal(body, &autoGen); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	// Extract values

	// array of type User with len 55
	users := make([]User, 0, 55)
	for _, userEntry := range autoGen.Data.XdtAPIV1FbsearchTopsearchConnection.Users {
		user := User{
			Fullname:      userEntry.User.Fullname,
			Username:      userEntry.User.Username,
			ID:            userEntry.User.ID,
			ProfilePicURL: "/proxy?url=" + url.QueryEscape(userEntry.User.ProfilePicURL),
		}
		users = append(users, user)
	}

	return users, nil
}

// returns -> list of usermedia, an offset string num and nil
func Timeline(endCursor, username string, client *http.Client) (PageData, error) {
	fmt.Printf("[+] Running Timeline(%s):\n", username)
	var data *strings.Reader
	if endCursor == "" {
		fmt.Printf("[+] First Page: %s", endCursor)
		// ADD DATA FOR NEXT PAGE HERE: takes Username
		// data = strings.NewReader()
	} else {
		fmt.Printf("[+] Next Page: %s", endCursor)
		// ADD DATA FOR NEXT PAGE HERE: takes encCursor + Username
		// data = strings.NewReader()
	}

	req, err := http.NewRequest("POST", "https://www.instagram.com/graphql/query", data)
	if err != nil {
		log.Fatal(err)
	}

	// ADD HEADERS HERE

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(`[+] response:  ` + string(body[:60]))

	// JSON parse
	var timelineAutoGen TimelineFirstJson
	if err := json.Unmarshal(body, &timelineAutoGen); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	println(`[+] Edges: `, len(timelineAutoGen.Data.XdtAPIV1FeedUserTimelineGraphqlConnection.Edges))

	// Extract values from Edges

	// array of TimelineMedia type with len of 12 posts
	userMediaArray := make([]TimelineMedia, 0, 12)

	for _, edge := range timelineAutoGen.Data.XdtAPIV1FeedUserTimelineGraphqlConnection.Edges {
		capion := edge.Node.Caption.Text
		var imageLink string
		if len(edge.Node.ImageVersions2.Candidates) > 0 {
			imageLink = edge.Node.ImageVersions2.Candidates[0].URL // first one is Highest quality
		} else {
			imageLink = "/static/404.jpg"
		}
		var videoLink string
		if len(edge.Node.VideoVersions) > 0 {
			videoLink = edge.Node.VideoVersions[0].URL // first one is Highest quality
		} else {
			videoLink = ""
		}
		// add to TimelineMedia Dict
		media := TimelineMedia{
			Caption: capion,
			Image:   "/proxy?url=" + url.QueryEscape(imageLink),
			Video:   videoLink,
		}
		userMediaArray = append(userMediaArray, media) // append and grow original array
	}
	fmt.Printf("[+] capacity and len of Media\n%d %d", cap(userMediaArray), len(userMediaArray))
	end_cursor := timelineAutoGen.Data.XdtAPIV1FeedUserTimelineGraphqlConnection.PageInfo.EndCursor
	pageData := PageData{
		Media:     userMediaArray,
		EndCursor: end_cursor,
		Username:  username,
	}
	return pageData, nil
}
