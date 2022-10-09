package main

import (
	"strconv"

	"github.com/rl404/nagato/mal"
)

// Uncomment the one you want to try
// and see the result in the json file.
func example() {
	// exampleGetAnimeDetails()
	// exampleGetAnimeList()
	// exampleGetAnimeRanking()
	// exampleGetSeasonalAnime()
	// exampleGetSuggestedAnime()

	// exampleUpdateMyAnimeListStatus()
	// exampleDeleteMyAnimeListStatus()
	exampleGetUserAnimeList()

	// exampleGetMangaDetails()
	// exampleGetMangaList()
	// exampleGetMangaRanking()

	// exampleUpdateMyMangaListStatus()
	// exampleDeleteMyMangaListStatus()
	// exampleGetUserMangaList()

	// exampleGetUserInfo()

	// exampleGetForum()
	// exampleGetForumTopics()
	// exampleGetForumTopicDetails()

	// exampleGetCharacterDetails()

	// exampleGetPeopleDetails()
}

func exampleGetAnimeDetails() {
	id := 44511

	d, err := malClient.GetAnimeDetails(id,
		"alternative_titles",
		"start_date",
		"end_date",
		"synopsis",
		"mean",
		"rank",
		"popularity",
		"num_list_users",
		"num_scoring_users",
		"nsfw",
		"genres",
		"created_at",
		"updated_at",
		"media_type",
		"status",
		"my_list_status",
		"num_episodes",
		"start_season",
		"broadcast",
		"source",
		"average_episode_duration",
		"rating",
		"studios",
		"pictures",
		"background",
		"related_anime{synopsis}",
		"related_manga",
		"recommendations{synopsis}",
		"statistics",
		"num_favorites",
		"opening_themes",
		"ending_themes",
		"videos",
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-"+strconv.Itoa(id), d)
}

func exampleGetAnimeList() {
	d, err := malClient.GetAnimeList(
		mal.GetAnimeListParam{
			Query:  "black rock",
			Limit:  5,
			Offset: 0,
		},
		"alternative_titles",
		"start_date",
		"end_date",
		"synopsis",
		"mean",
		"rank",
		"popularity",
		"num_list_users",
		"num_scoring_users",
		"nsfw",
		"genres",
		"created_at",
		"updated_at",
		"media_type",
		"status",
		"my_list_status",
		"num_episodes",
		"start_season",
		"broadcast",
		"source",
		"average_episode_duration",
		"rating",
		"studios",
		"num_favorites",
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-list", d)
}

func exampleGetAnimeRanking() {
	d, err := malClient.GetAnimeRanking(
		mal.GetAnimeRankingParam{
			RankingType: "airing",
			Limit:       5,
			Offset:      0,
		},
		"alternative_titles",
		"start_date",
		"end_date",
		"synopsis",
		"mean",
		"rank",
		"popularity",
		"num_list_users",
		"num_scoring_users",
		"nsfw",
		"genres",
		"created_at",
		"updated_at",
		"media_type",
		"status",
		"my_list_status",
		"num_episodes",
		"start_season",
		"broadcast",
		"source",
		"average_episode_duration",
		"rating",
		"studios",
		"num_favorites",
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-ranking", d)
}

func exampleGetSeasonalAnime() {
	d, err := malClient.GetSeasonalAnime(
		mal.GetSeasonalAnimeParam{
			Year:   2021,
			Season: "fall",
			Sort:   "anime_num_list_users",
			Limit:  5,
			Offset: 0,
		},
		"alternative_titles",
		"start_date",
		"end_date",
		"synopsis",
		"mean",
		"rank",
		"popularity",
		"num_list_users",
		"num_scoring_users",
		"nsfw",
		"genres",
		"created_at",
		"updated_at",
		"media_type",
		"status",
		"my_list_status",
		"num_episodes",
		"start_season",
		"broadcast",
		"source",
		"average_episode_duration",
		"rating",
		"studios",
		"num_favorites",
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-season", d)
}

func exampleGetSuggestedAnime() {
	d, err := malClient.GetSuggestedAnime(
		mal.GetSuggestedAnimeParam{
			Limit:  5,
			Offset: 0,
		},
		"alternative_titles",
		"start_date",
		"end_date",
		"synopsis",
		"mean",
		"rank",
		"popularity",
		"num_list_users",
		"num_scoring_users",
		"nsfw",
		"genres",
		"created_at",
		"updated_at",
		"media_type",
		"status",
		"my_list_status",
		"num_episodes",
		"start_season",
		"broadcast",
		"source",
		"average_episode_duration",
		"rating",
		"studios",
		"num_favorites",
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-suggestion", d)
}

func exampleGetMangaDetails() {
	id := 119161

	d, err := malClient.GetMangaDetails(id,
		"alternative_titles",
		"start_date",
		"end_date",
		"synopsis",
		"mean",
		"rank",
		"popularity",
		"num_list_users",
		"num_scoring_users",
		"nsfw",
		"genres",
		"created_at",
		"updated_at",
		"media_type",
		"status",
		"my_list_status",
		"num_volumes",
		"num_chapters",
		"authors{first_name,last_name}",
		"pictures",
		"background",
		"related_anime",
		"related_manga{synopsis}",
		"recommendations{synopsis}",
		"serialization",
		"num_favorites",
	)
	if err != nil {
		panic(err)
	}

	toJson("manga-"+strconv.Itoa(id), d)
}

func exampleGetMangaList() {
	d, err := malClient.GetMangaList(mal.GetMangaListParam{
		Query:  "naruto",
		Nsfw:   true,
		Limit:  5,
		Offset: 0,
	},
		"alternative_titles",
		"start_date",
		"end_date",
		"synopsis",
		"mean",
		"rank",
		"popularity",
		"num_list_users",
		"num_scoring_users",
		"nsfw",
		"genres",
		"created_at",
		"updated_at",
		"media_type",
		"status",
		"my_list_status",
		"num_volumes",
		"num_chapters",
		"authors{first_name,last_name}",
		"num_favorites",
	)
	if err != nil {
		panic(err)
	}

	toJson("manga-list", d)
}

func exampleGetMangaRanking() {
	d, err := malClient.GetMangaRanking(mal.GetMangaRankingParam{
		RankingType: "manga",
		Nsfw:        true,
		Limit:       5,
		Offset:      0,
	},
		"alternative_titles",
		"start_date",
		"end_date",
		"synopsis",
		"mean",
		"rank",
		"popularity",
		"num_list_users",
		"num_scoring_users",
		"nsfw",
		"genres",
		"created_at",
		"updated_at",
		"media_type",
		"status",
		"my_list_status",
		"num_volumes",
		"num_chapters",
		"authors{first_name,last_name}",
		"num_favorites",
	)
	if err != nil {
		panic(err)
	}

	toJson("manga-ranking", d)
}

func exampleUpdateMyAnimeListStatus() {
	d, err := malClient.UpdateMyAnimeListStatus(mal.UpdateMyAnimeListStatusParam{
		ID:                 44511,
		Status:             "plan_to_watch",
		NumWatchedEpisodes: 1,
		IsRewatching:       true,
		Score:              8,
		Priority:           1,
		NumTimesRewatched:  2,
		RewatchValue:       3,
		Tags:               "tags1,tags2",
		Comments:           "comments",
		StartDate:          "2020-10-12",
		FinishDate:         "2020-11-12",
	})
	if err != nil {
		panic(err)
	}

	toJson("update-anime-status", d)
}

func exampleDeleteMyAnimeListStatus() {
	err := malClient.DeleteMyAnimeListStatus(44511)
	if err != nil {
		panic(err)
	}
}

func exampleGetUserAnimeList() {
	username := "rl404"

	d, err := malClient.GetUserAnimeList(
		mal.GetUserAnimeListParam{
			Username: username,
			Status:   "watching",
			Nsfw:     true,
			Sort:     "",
			Limit:    5,
			Offset:   0,
		},
		"list_status{tags}",
		"alternative_titles",
		"start_date",
		"end_date",
		"synopsis",
		"mean",
		"rank",
		"popularity",
		"num_list_users",
		"num_scoring_users",
		"nsfw",
		"genres",
		"created_at",
		"updated_at",
		"media_type",
		"status",
		"my_list_status",
		"num_episodes",
		"start_season",
		"broadcast",
		"source",
		"average_episode_duration",
		"rating",
		"studios",
		"num_favorites",
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-list-"+username, d)
}

func exampleUpdateMyMangaListStatus() {
	d, err := malClient.UpdateMyMangaListStatus(mal.UpdateMyMangaListStatusParam{
		ID:              83903,
		Status:          "plan_to_read",
		NumVolumesRead:  1,
		NumChaptersRead: 3,
		IsRereading:     true,
		Score:           8,
		Priority:        1,
		NumTimesReread:  2,
		RereadValue:     3,
		Tags:            "tags1,tags2",
		Comments:        "comments",
		StartDate:       "2020-10-12",
		FinishDate:      "2020-11-12",
	})
	if err != nil {
		panic(err)
	}

	toJson("update-manga-status", d)
}

func exampleDeleteMyMangaListStatus() {
	err := malClient.DeleteMyMangaListStatus(83903)
	if err != nil {
		panic(err)
	}
}

func exampleGetUserMangaList() {
	username := "rl404"

	d, err := malClient.GetUserMangaList(
		mal.GetUserMangaListParam{
			Username: username,
			Status:   "",
			Nsfw:     true,
			Sort:     "",
			Limit:    5,
			Offset:   0,
		},
		"list_status{tags}",
		"alternative_titles",
		"start_date",
		"end_date",
		"synopsis",
		"mean",
		"rank",
		"popularity",
		"num_list_users",
		"num_scoring_users",
		"nsfw",
		"genres",
		"created_at",
		"updated_at",
		"media_type",
		"status",
		"my_list_status",
		"num_volumes",
		"num_chapters",
		"authors{first_name,last_name}",
		"num_favorites",
	)
	if err != nil {
		panic(err)
	}

	toJson("manga-list-"+username, d)
}

func exampleGetUserInfo() {
	username := "@me"

	d, err := malClient.GetUserInfo(username,
		"time_zone",
		"anime_statistics",
	)
	if err != nil {
		panic(err)
	}

	toJson("user-info-"+username, d)
}

func exampleGetForum() {
	d, err := malClient.GetForumBoards()
	if err != nil {
		panic(err)
	}

	toJson("forum-boards", d)
}

func exampleGetForumTopics() {
	d, err := malClient.GetForumTopics(mal.GetForumTopicsParam{
		BoardID:       0,
		SubboardID:    2,
		Query:         "anime",
		TopicUsername: "Kineta",
		Username:      "",
		Sort:          "recent",
		Limit:         5,
		Offset:        0,
	})
	if err != nil {
		panic(err)
	}

	toJson("forum-topics", d)
}

func exampleGetForumTopicDetails() {
	id := 2048122

	d, err := malClient.GetForumTopicDetails(mal.GetForumTopicDetailsParam{
		ID:     id,
		Limit:  5,
		Offset: 0,
	})
	if err != nil {
		panic(err)
	}

	toJson("forum-topic-"+strconv.Itoa(id), d)
}

func exampleGetCharacterDetails() {
	id := 170329

	d, err := malClient.GetCharacterDetails(id,
		"first_name",
		"last_name",
		"alternative_name",
		"main_picture",
		"biography",
		"num_favorites",
		"animeography",
	)
	if err != nil {
		panic(err)
	}

	toJson("character-"+strconv.Itoa(id), d)
}

func exampleGetPeopleDetails() {
	id := 869

	d, err := malClient.GetPeopleDetails(id,
		"first_name",
		"last_name",
		"alternative_names",
		"main_picture",
		"birthday",
		"website_url",
		"num_favorites",
		"more",
	)

	if err != nil {
		panic(err)
	}

	toJson("people-"+strconv.Itoa(id), d)
}
