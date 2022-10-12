package main

import (
	"strconv"

	"github.com/rl404/nagato"
)

func example() {
	// exampleGetAnimeDetails()
	// exampleGetAnimeList()
	// exampleGetAnimeRanking()
	// exampleGetSeasonalAnime()
	// exampleGetSuggestedAnime()

	// exampleGetMangaDetails()
	// exampleGetMangaList()
	// exampleGetMangaRanking()

	// exampleGetUserAnimeList()
	exampleGetUserMangaList()
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
