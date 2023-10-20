package main

import (
	"strconv"

	"github.com/rl404/nagato"
)

func example() {
	exampleGetAnimeDetails()
	// exampleGetAnimeList()
	// exampleGetAnimeRanking()
	// exampleGetSeasonalAnime()
	// exampleGetSuggestedAnime()

	// exampleGetMangaDetails()
	// exampleGetMangaList()
	// exampleGetMangaRanking()

	// exampleGetUserAnimeList()
	// exampleUpdateMyAnimeListStatus()
	// exampleDeleteMyAnimeListStatus()

	// exampleGetUserMangaList()
	// exampleUpdateMyMangaListStatus()
	// exampleDeleteMyMangaListStatus()

	// exampleGetUserInfo()

	// exampleGetForumBoards()
	// exampleGetForumTopics()
	// exampleGetForumTopicDetails()
}

func exampleGetAnimeDetails() {
	id := 1

	d, _, err := nagatoClient.GetAnimeDetails(id,
		nagato.AnimeFieldAlternativeTitles,
		nagato.AnimeFieldStartDate,
		nagato.AnimeFieldEndDate,
		nagato.AnimeFieldSynopsis,
		nagato.AnimeFieldMean,
		nagato.AnimeFieldRank,
		nagato.AnimeFieldPopularity,
		nagato.AnimeFieldNumListUsers,
		nagato.AnimeFieldNumScoringUsers,
		nagato.AnimeFieldNSFW,
		nagato.AnimeFieldGenres,
		nagato.AnimeFieldMediaType,
		nagato.AnimeFieldStatus,
		nagato.AnimeFieldMyListStatus,
		nagato.AnimeFieldNumEpisodes,
		nagato.AnimeFieldStartSeason,
		nagato.AnimeFieldBroadcast,
		nagato.AnimeFieldSource,
		nagato.AnimeFieldAverageEpisodeDuration,
		nagato.AnimeFieldRating,
		nagato.AnimeFieldStudios,
		nagato.AnimeFieldPictures,
		nagato.AnimeFieldBackground,
		nagato.AnimeFieldRelatedAnime(nagato.AnimeFieldAlternativeTitles),
		nagato.AnimeFieldRecommendations(nagato.AnimeFieldAlternativeTitles),
		nagato.AnimeFieldStatistics,
		nagato.AnimeFieldNumFavorites,
		nagato.AnimeFieldOpeningThemes,
		nagato.AnimeFieldEndingThemes,
		nagato.AnimeFieldVideos,
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-"+strconv.Itoa(id), d)
}

func exampleGetAnimeList() {
	d, _, err := nagatoClient.GetAnimeList(nagato.GetAnimeListParam{
		Query:  "gurren",
		NSFW:   true,
		Limit:  5,
		Offset: 0,
	},
		nagato.AnimeFieldAlternativeTitles,
		nagato.AnimeFieldStartDate,
		nagato.AnimeFieldEndDate,
		nagato.AnimeFieldSynopsis,
		nagato.AnimeFieldMean,
		nagato.AnimeFieldRank,
		nagato.AnimeFieldPopularity,
		nagato.AnimeFieldNumListUsers,
		nagato.AnimeFieldNumScoringUsers,
		nagato.AnimeFieldNSFW,
		nagato.AnimeFieldGenres,
		nagato.AnimeFieldMediaType,
		nagato.AnimeFieldStatus,
		nagato.AnimeFieldMyListStatus,
		nagato.AnimeFieldNumEpisodes,
		nagato.AnimeFieldStartSeason,
		nagato.AnimeFieldBroadcast,
		nagato.AnimeFieldSource,
		nagato.AnimeFieldAverageEpisodeDuration,
		nagato.AnimeFieldRating,
		nagato.AnimeFieldStudios,
		nagato.AnimeFieldNumFavorites,
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-list", d)
}

func exampleGetAnimeRanking() {
	d, _, err := nagatoClient.GetAnimeRanking(nagato.GetAnimeRankingParam{
		RankingType: nagato.RankingAll,
		NSFW:        true,
		Limit:       5,
		Offset:      0,
	},
		nagato.AnimeFieldAlternativeTitles,
		nagato.AnimeFieldStartDate,
		nagato.AnimeFieldEndDate,
		nagato.AnimeFieldSynopsis,
		nagato.AnimeFieldMean,
		nagato.AnimeFieldRank,
		nagato.AnimeFieldPopularity,
		nagato.AnimeFieldNumListUsers,
		nagato.AnimeFieldNumScoringUsers,
		nagato.AnimeFieldNSFW,
		nagato.AnimeFieldGenres,
		nagato.AnimeFieldMediaType,
		nagato.AnimeFieldStatus,
		nagato.AnimeFieldMyListStatus,
		nagato.AnimeFieldNumEpisodes,
		nagato.AnimeFieldStartSeason,
		nagato.AnimeFieldBroadcast,
		nagato.AnimeFieldSource,
		nagato.AnimeFieldAverageEpisodeDuration,
		nagato.AnimeFieldRating,
		nagato.AnimeFieldStudios,
		nagato.AnimeFieldNumFavorites,
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-ranking", d)
}

func exampleGetSeasonalAnime() {
	d, _, err := nagatoClient.GetSeasonalAnime(nagato.GetSeasonalAnimeParam{
		Year:   2020,
		Season: nagato.SeasonFall,
		NSFW:   true,
		Sort:   nagato.SeasonalAnimeSortByScore,
		Limit:  5,
		Offset: 0,
	},
		nagato.AnimeFieldAlternativeTitles,
		nagato.AnimeFieldStartDate,
		nagato.AnimeFieldEndDate,
		nagato.AnimeFieldSynopsis,
		nagato.AnimeFieldMean,
		nagato.AnimeFieldRank,
		nagato.AnimeFieldPopularity,
		nagato.AnimeFieldNumListUsers,
		nagato.AnimeFieldNumScoringUsers,
		nagato.AnimeFieldNSFW,
		nagato.AnimeFieldGenres,
		nagato.AnimeFieldMediaType,
		nagato.AnimeFieldStatus,
		nagato.AnimeFieldMyListStatus,
		nagato.AnimeFieldNumEpisodes,
		nagato.AnimeFieldStartSeason,
		nagato.AnimeFieldBroadcast,
		nagato.AnimeFieldSource,
		nagato.AnimeFieldAverageEpisodeDuration,
		nagato.AnimeFieldRating,
		nagato.AnimeFieldStudios,
		nagato.AnimeFieldNumFavorites,
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-seasonal", d)
}

func exampleGetSuggestedAnime() {
	d, _, err := nagatoClient.GetSuggestedAnime(nagato.GetSuggestedAnimeParam{
		NSFW:   true,
		Limit:  5,
		Offset: 0,
	},
		nagato.AnimeFieldAlternativeTitles,
		nagato.AnimeFieldStartDate,
		nagato.AnimeFieldEndDate,
		nagato.AnimeFieldSynopsis,
		nagato.AnimeFieldMean,
		nagato.AnimeFieldRank,
		nagato.AnimeFieldPopularity,
		nagato.AnimeFieldNumListUsers,
		nagato.AnimeFieldNumScoringUsers,
		nagato.AnimeFieldNSFW,
		nagato.AnimeFieldGenres,
		nagato.AnimeFieldMediaType,
		nagato.AnimeFieldStatus,
		nagato.AnimeFieldMyListStatus,
		nagato.AnimeFieldNumEpisodes,
		nagato.AnimeFieldStartSeason,
		nagato.AnimeFieldBroadcast,
		nagato.AnimeFieldSource,
		nagato.AnimeFieldAverageEpisodeDuration,
		nagato.AnimeFieldRating,
		nagato.AnimeFieldStudios,
		nagato.AnimeFieldNumFavorites,
	)
	if err != nil {
		panic(err)
	}

	toJson("anime-suggestion", d)
}

func exampleGetMangaDetails() {
	id := 1

	d, _, err := nagatoClient.GetMangaDetails(id,
		nagato.MangaFieldAlternativeTitles,
		nagato.MangaFieldStartDate,
		nagato.MangaFieldEndDate,
		nagato.MangaFieldSynopsis,
		nagato.MangaFieldMean,
		nagato.MangaFieldRank,
		nagato.MangaFieldPopularity,
		nagato.MangaFieldNumListUsers,
		nagato.MangaFieldNumScoringUsers,
		nagato.MangaFieldNSFW,
		nagato.MangaFieldGenres,
		nagato.MangaFieldMediaType,
		nagato.MangaFieldStatus,
		nagato.MangaFieldMyListStatus,
		nagato.MangaFieldNumVolumes,
		nagato.MangaFieldNumChapters,
		nagato.MangaFieldAuthors,
		nagato.MangaFieldPictures,
		nagato.MangaFieldBackground,
		nagato.MangaFieldRelatedManga(nagato.MangaFieldAlternativeTitles),
		nagato.MangaFieldRecommendations(nagato.MangaFieldAlternativeTitles),
		nagato.MangaFieldSerialization,
		nagato.MangaFieldNumFavorites,
	)
	if err != nil {
		panic(err)
	}

	toJson("manga-"+strconv.Itoa(id), d)
}

func exampleGetMangaList() {
	d, _, err := nagatoClient.GetMangaList(nagato.GetMangaListParam{
		Query:  "naruto",
		NSFW:   true,
		Limit:  5,
		Offset: 0,
	},
		nagato.MangaFieldAlternativeTitles,
		nagato.MangaFieldStartDate,
		nagato.MangaFieldEndDate,
		nagato.MangaFieldSynopsis,
		nagato.MangaFieldMean,
		nagato.MangaFieldRank,
		nagato.MangaFieldPopularity,
		nagato.MangaFieldNumListUsers,
		nagato.MangaFieldNumScoringUsers,
		nagato.MangaFieldNSFW,
		nagato.MangaFieldGenres,
		nagato.MangaFieldMediaType,
		nagato.MangaFieldStatus,
		nagato.MangaFieldMyListStatus,
		nagato.MangaFieldNumVolumes,
		nagato.MangaFieldNumChapters,
		nagato.MangaFieldAuthors,
		nagato.MangaFieldNumFavorites,
	)
	if err != nil {
		panic(err)
	}

	toJson("manga-list", d)
}

func exampleGetMangaRanking() {
	d, _, err := nagatoClient.GetMangaRanking(nagato.GetMangaRankingParam{
		RankingType: nagato.RankingAll,
		NSFW:        true,
		Limit:       5,
		Offset:      0,
	},
		nagato.MangaFieldAlternativeTitles,
		nagato.MangaFieldStartDate,
		nagato.MangaFieldEndDate,
		nagato.MangaFieldSynopsis,
		nagato.MangaFieldMean,
		nagato.MangaFieldRank,
		nagato.MangaFieldPopularity,
		nagato.MangaFieldNumListUsers,
		nagato.MangaFieldNumScoringUsers,
		nagato.MangaFieldNSFW,
		nagato.MangaFieldGenres,
		nagato.MangaFieldMediaType,
		nagato.MangaFieldStatus,
		nagato.MangaFieldMyListStatus,
		nagato.MangaFieldNumVolumes,
		nagato.MangaFieldNumChapters,
		nagato.MangaFieldAuthors,
		nagato.MangaFieldNumFavorites,
	)
	if err != nil {
		panic(err)
	}

	toJson("manga-ranking", d)
}

func exampleGetUserAnimeList() {
	d, _, err := nagatoClient.GetUserAnimeList(nagato.GetUserAnimeListParam{
		Username: "rl404",
		Status:   nagato.UserAnimeStatusCompleted,
		NSFW:     true,
		Sort:     nagato.UserAnimeSortScore,
		Limit:    5,
		Offset:   0,
	},
		nagato.AnimeFieldAlternativeTitles,
		nagato.AnimeFieldStartDate,
		nagato.AnimeFieldEndDate,
		nagato.AnimeFieldSynopsis,
		nagato.AnimeFieldMean,
		nagato.AnimeFieldRank,
		nagato.AnimeFieldPopularity,
		nagato.AnimeFieldNumListUsers,
		nagato.AnimeFieldNumScoringUsers,
		nagato.AnimeFieldNSFW,
		nagato.AnimeFieldGenres,
		nagato.AnimeFieldMediaType,
		nagato.AnimeFieldStatus,
		nagato.AnimeFieldMyListStatus,
		nagato.AnimeFieldNumEpisodes,
		nagato.AnimeFieldStartSeason,
		nagato.AnimeFieldBroadcast,
		nagato.AnimeFieldSource,
		nagato.AnimeFieldAverageEpisodeDuration,
		nagato.AnimeFieldRating,
		nagato.AnimeFieldStudios,
		nagato.AnimeFieldNumFavorites,
		nagato.AnimeFieldUserStatus(nagato.UserAnimeTags),
	)
	if err != nil {
		panic(err)
	}

	toJson("user-anime", d)
}

func exampleUpdateMyAnimeListStatus() {
	d, _, err := nagatoClient.UpdateMyAnimeListStatus(nagato.UpdateMyAnimeListStatusParam{
		ID:             1,
		Status:         nagato.UserAnimeStatusPlanToWatch,
		IsRewatching:   true,
		Episode:        1,
		Score:          8,
		Priority:       1,
		RewatchedTimes: 2,
		RewatchValue:   3,
		Tags:           []string{"tag1", "tag2"},
		Comment:        "comments",
		StartDate: nagato.Date{
			Year:  2022,
			Month: 10,
			Day:   31,
		},
		FinishDate: nagato.Date{
			Year:  2021,
			Month: 10,
			Day:   1,
		},
	})
	if err != nil {
		panic(err)
	}

	toJson("update-anime-status", d)
}

func exampleDeleteMyAnimeListStatus() {
	_, err := nagatoClient.DeleteMyAnimeListStatus(1)
	if err != nil {
		panic(err)
	}
}

func exampleGetUserMangaList() {
	d, _, err := nagatoClient.GetUserMangaList(nagato.GetUserMangaListParam{
		Username: "rl404",
		Status:   nagato.UserMangaStatusCompleted,
		NSFW:     true,
		Sort:     nagato.UserMangaSortScore,
		Limit:    5,
		Offset:   0,
	},
		nagato.MangaFieldAlternativeTitles,
		nagato.MangaFieldStartDate,
		nagato.MangaFieldEndDate,
		nagato.MangaFieldSynopsis,
		nagato.MangaFieldMean,
		nagato.MangaFieldRank,
		nagato.MangaFieldPopularity,
		nagato.MangaFieldNumListUsers,
		nagato.MangaFieldNumScoringUsers,
		nagato.MangaFieldNSFW,
		nagato.MangaFieldGenres,
		nagato.MangaFieldMediaType,
		nagato.MangaFieldStatus,
		nagato.MangaFieldMyListStatus,
		nagato.MangaFieldNumVolumes,
		nagato.MangaFieldNumChapters,
		nagato.MangaFieldAuthors,
		nagato.MangaFieldNumFavorites,
		nagato.MangaFieldUserStatus(nagato.UserMangaTags),
	)
	if err != nil {
		panic(err)
	}

	toJson("user-manga", d)
}

func exampleUpdateMyMangaListStatus() {
	d, _, err := nagatoClient.UpdateMyMangaListStatus(nagato.UpdateMyMangaListStatusParam{
		ID:          1706,
		Status:      nagato.UserMangaStatusPlanToRead,
		IsRereading: true,
		Volume:      1,
		Chapter:     2,
		Score:       8,
		Priority:    1,
		RereadTimes: 2,
		RereadValue: 3,
		Tags:        []string{"tag1", "tag2"},
		Comment:     "comments",
		StartDate: nagato.Date{
			Year:  2022,
			Month: 10,
			Day:   31,
		},
		FinishDate: nagato.Date{
			Year:  2021,
			Month: 10,
			Day:   1,
		},
	})
	if err != nil {
		panic(err)
	}

	toJson("update-anime-status", d)
}

func exampleDeleteMyMangaListStatus() {
	_, err := nagatoClient.DeleteMyMangaListStatus(1706)
	if err != nil {
		panic(err)
	}
}

func exampleGetUserInfo() {
	d, _, err := nagatoClient.GetUserInfo("",
		nagato.UserFieldAnimeStatistics,
		nagato.UserFieldTimeZone,
	)
	if err != nil {
		panic(err)
	}

	toJson("user", d)
}

func exampleGetForumBoards() {
	d, _, err := nagatoClient.GetForumBoards()
	if err != nil {
		panic(err)
	}

	toJson("forum-boards", d)
}

func exampleGetForumTopics() {
	d, _, err := nagatoClient.GetForumTopics(nagato.GetForumTopicsParam{
		BoardID:       0,
		SubboardID:    0,
		Query:         "",
		TopicUsername: "",
		Username:      "",
		Sort:          nagato.ForumTopicSortRecent,
		Limit:         5,
		Offset:        0,
	})
	if err != nil {
		panic(err)
	}

	toJson("forum-topics", d)
}

func exampleGetForumTopicDetails() {
	d, _, err := nagatoClient.GetForumTopicDetails(nagato.GetForumTopicDetailsParam{
		ID:     2053204,
		Limit:  5,
		Offset: 6,
	})
	if err != nil {
		panic(err)
	}

	toJson("forum-topic", d)
}
