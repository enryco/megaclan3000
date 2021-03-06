package steamclient

// GameStats holds the players stats data from the steam API
// endpoint UserStatsForGame
type GameStats struct {
	SteamID                                   string `db:"steamid"`
	GILessonCsgoInstrExplainInspect           string `db:"gi_lesson_csgo_instr_explain_inspect"`
	GILessonBombSitesA                        string `db:"gi_lesson_bomb_sites_a"`
	GILessonBombSitesB                        string `db:"gi_lesson_bomb_sites_b"`
	GILessonCsgoCycleWeaponsKb                string `db:"gi_lesson_csgo_cycle_weapons_kb"`
	GILessonCsgoHostageLeadToHrz              string `db:"gi_lesson_csgo_hostage_lead_to_hrz"`
	GILessonCsgoInstrExplainBombCarrier       string `db:"gi_lesson_csgo_instr_explain_bomb_carrier"`
	GILessonCsgoInstrExplainBuyarmor          string `db:"gi_lesson_csgo_instr_explain_buyarmor"`
	GILessonCsgoInstrExplainBuymenu           string `db:"gi_lesson_csgo_instr_explain_buymenu"`
	GILessonCsgoInstrExplainFollowBomber      string `db:"gi_lesson_csgo_instr_explain_follow_bomber"`
	GILessonCsgoInstrExplainPickupBomb        string `db:"gi_lesson_csgo_instr_explain_pickup_bomb"`
	GILessonCsgoInstrExplainPlantBomb         string `db:"gi_lesson_csgo_instr_explain_plant_bomb"`
	GILessonCsgoInstrExplainPreventBombPickup string `db:"gi_lesson_csgo_instr_explain_prevent_bomb_pickup"`
	GILessonCsgoInstrExplainReload            string `db:"gi_lesson_csgo_instr_explain_reload"`
	GILessonCsgoInstrExplainZoom              string `db:"gi_lesson_csgo_instr_explain_zoom"`
	GILessonCsgoInstrRescueZone               string `db:"gi_lesson_csgo_instr_rescue_zone"`
	GILessonDefusePlantedBomb                 string `db:"gi_lesson_defuse_planted_bomb"`
	GILessonFindPlantedBomb                   string `db:"gi_lesson_find_planted_bomb"`
	GILessonTrExplainPlantBomb                string `db:"gi_lesson_tr_explain_plant_bomb"`
	GILessonVersionNumber                     string `db:"gi_lesson_version_number"`
	LastMatchContributionScore                string `db:"last_match_contribution_score"`
	LastMatchCtWins                           string `db:"last_match_ct_wins"`
	LastMatchDamage                           string `db:"last_match_damage"`
	LastMatchDeaths                           string `db:"last_match_deaths"`
	LastMatchDominations                      string `db:"last_match_dominations"`
	LastMatchFavweaponHits                    string `db:"last_match_favweapon_hits"`
	LastMatchFavweaponID                      string `db:"last_match_favweapon_id"`
	LastMatchFavweaponKills                   string `db:"last_match_favweapon_kills"`
	LastMatchFavweaponShots                   string `db:"last_match_favweapon_shots"`
	LastMatchGgContributionScore              string `db:"last_match_gg_contribution_score"`
	LastMatchKills                            string `db:"last_match_kills"`
	LastMatchMaxPlayers                       string `db:"last_match_max_players"`
	LastMatchMoneySpent                       string `db:"last_match_money_spent"`
	LastMatchMvps                             string `db:"last_match_mvps"`
	LastMatchRevenges                         string `db:"last_match_revenges"`
	LastMatchRounds                           string `db:"last_match_rounds"`
	LastMatchTWins                            string `db:"last_match_t_wins"`
	LastMatchWins                             string `db:"last_match_wins"`
	SteamStatMatchwinscomp                    string `db:"steam_stat_matchwinscomp"`
	SteamStatSurvivedz                        string `db:"steam_stat_survivedz"`
	SteamStatXpearnedgames                    string `db:"steam_stat_xpearnedgames"`
	TotalBrokenWindows                        string `db:"total_broken_windows"`
	TotalContributionScore                    string `db:"total_contribution_score"`
	TotalDamageDone                           string `db:"total_damage_done"`
	TotalDeaths                               string `db:"total_deaths"`
	TotalDefusedBombs                         string `db:"total_defused_bombs"`
	TotalDominationOverkills                  string `db:"total_domination_overkills"`
	TotalDominations                          string `db:"total_dominations"`
	TotalGgMatchesPlayed                      string `db:"total_gg_matches_played"`
	TotalGgMatchesWon                         string `db:"total_gg_matches_won"`
	TotalGunGameContributionScore             string `db:"total_gun_game_contribution_score"`
	TotalGunGameRoundsPlayed                  string `db:"total_gun_game_rounds_played"`
	TotalGunGameRoundsWon                     string `db:"total_gun_game_rounds_won"`
	TotalHitsAk47                             string `db:"total_hits_ak47"`
	TotalHitsAug                              string `db:"total_hits_aug"`
	TotalHitsAwp                              string `db:"total_hits_awp"`
	TotalHitsBizon                            string `db:"total_hits_bizon"`
	TotalHitsDeagle                           string `db:"total_hits_deagle"`
	TotalHitsElite                            string `db:"total_hits_elite"`
	TotalHitsFamas                            string `db:"total_hits_famas"`
	TotalHitsFiveseven                        string `db:"total_hits_fiveseven"`
	TotalHitsG3sg1                            string `db:"total_hits_g3sg1"`
	TotalHitsGalilar                          string `db:"total_hits_galilar"`
	TotalHitsGlock                            string `db:"total_hits_glock"`
	TotalHitsHkp2000                          string `db:"total_hits_hkp2000"`
	TotalHitsM249                             string `db:"total_hits_m249"`
	TotalHitsM4a1                             string `db:"total_hits_m4a1"`
	TotalHitsMac10                            string `db:"total_hits_mac10"`
	TotalHitsMag7                             string `db:"total_hits_mag7"`
	TotalHitsMp7                              string `db:"total_hits_mp7"`
	TotalHitsMp9                              string `db:"total_hits_mp9"`
	TotalHitsNegev                            string `db:"total_hits_negev"`
	TotalHitsNova                             string `db:"total_hits_nova"`
	TotalHitsP250                             string `db:"total_hits_p250"`
	TotalHitsP90                              string `db:"total_hits_p90"`
	TotalHitsSawedoff                         string `db:"total_hits_sawedoff"`
	TotalHitsScar20                           string `db:"total_hits_scar20"`
	TotalHitsSg556                            string `db:"total_hits_sg556"`
	TotalHitsSsg08                            string `db:"total_hits_ssg08"`
	TotalHitsTec9                             string `db:"total_hits_tec9"`
	TotalHitsUmp45                            string `db:"total_hits_ump45"`
	TotalHitsXm1014                           string `db:"total_hits_xm1014"`
	TotalKills                                string `db:"total_kills"`
	TotalKillsAgainstZoomedSniper             string `db:"total_kills_against_zoomed_sniper"`
	TotalKillsAk47                            string `db:"total_kills_ak47"`
	TotalKillsAug                             string `db:"total_kills_aug"`
	TotalKillsAwp                             string `db:"total_kills_awp"`
	TotalKillsBizon                           string `db:"total_kills_bizon"`
	TotalKillsDeagle                          string `db:"total_kills_deagle"`
	TotalKillsElite                           string `db:"total_kills_elite"`
	TotalKillsEnemyBlinded                    string `db:"total_kills_enemy_blinded"`
	TotalKillsEnemyWeapon                     string `db:"total_kills_enemy_weapon"`
	TotalKillsFamas                           string `db:"total_kills_famas"`
	TotalKillsFiveseven                       string `db:"total_kills_fiveseven"`
	TotalKillsG3sg1                           string `db:"total_kills_g3sg1"`
	TotalKillsGalilar                         string `db:"total_kills_galilar"`
	TotalKillsGlock                           string `db:"total_kills_glock"`
	TotalKillsHeadshot                        string `db:"total_kills_headshot"`
	TotalKillsHegrenade                       string `db:"total_kills_hegrenade"`
	TotalKillsHkp2000                         string `db:"total_kills_hkp2000"`
	TotalKillsKnife                           string `db:"total_kills_knife"`
	TotalKillsKnifeFight                      string `db:"total_kills_knife_fight"`
	TotalKillsM249                            string `db:"total_kills_m249"`
	TotalKillsM4a1                            string `db:"total_kills_m4a1"`
	TotalKillsMac10                           string `db:"total_kills_mac10"`
	TotalKillsMag7                            string `db:"total_kills_mag7"`
	TotalKillsMolotov                         string `db:"total_kills_molotov"`
	TotalKillsMp7                             string `db:"total_kills_mp7"`
	TotalKillsMp9                             string `db:"total_kills_mp9"`
	TotalKillsNegev                           string `db:"total_kills_negev"`
	TotalKillsNova                            string `db:"total_kills_nova"`
	TotalKillsP250                            string `db:"total_kills_p250"`
	TotalKillsP90                             string `db:"total_kills_p90"`
	TotalKillsSawedoff                        string `db:"total_kills_sawedoff"`
	TotalKillsScar20                          string `db:"total_kills_scar20"`
	TotalKillsSg556                           string `db:"total_kills_sg556"`
	TotalKillsSsg08                           string `db:"total_kills_ssg08"`
	TotalKillsTaser                           string `db:"total_kills_taser"`
	TotalKillsTec9                            string `db:"total_kills_tec9"`
	TotalKillsUmp45                           string `db:"total_kills_ump45"`
	TotalKillsXm1014                          string `db:"total_kills_xm1014"`
	TotalMatchesPlayed                        string `db:"total_matches_played"`
	TotalMatchesWon                           string `db:"total_matches_won"`
	TotalMatchesWonBaggage                    string `db:"total_matches_won_baggage"`
	TotalMatchesWonBank                       string `db:"total_matches_won_bank"`
	TotalMatchesWonLake                       string `db:"total_matches_won_lake"`
	TotalMatchesWonSafehouse                  string `db:"total_matches_won_safehouse"`
	TotalMatchesWonShoots                     string `db:"total_matches_won_shoots"`
	TotalMatchesWonStmarc                     string `db:"total_matches_won_stmarc"`
	TotalMatchesWonSugarcane                  string `db:"total_matches_won_sugarcane"`
	TotalMatchesWonTrain                      string `db:"total_matches_won_train"`
	TotalMoneyEarned                          string `db:"total_money_earned"`
	TotalMvps                                 string `db:"total_mvps"`
	TotalPlantedBombs                         string `db:"total_planted_bombs"`
	TotalProgressiveMatchesWon                string `db:"total_progressive_matches_won"`
	TotalRescuedHostages                      string `db:"total_rescued_hostages"`
	TotalRevenges                             string `db:"total_revenges"`
	TotalRoundsMapArBaggage                   string `db:"total_rounds_map_ar_baggage"`
	TotalRoundsMapArMonastery                 string `db:"total_rounds_map_ar_monastery"`
	TotalRoundsMapArShoots                    string `db:"total_rounds_map_ar_shoots"`
	TotalRoundsMapCsAssault                   string `db:"total_rounds_map_cs_assault"`
	TotalRoundsMapCsItaly                     string `db:"total_rounds_map_cs_italy"`
	TotalRoundsMapCsMilitia                   string `db:"total_rounds_map_cs_militia"`
	TotalRoundsMapCsOffice                    string `db:"total_rounds_map_cs_office"`
	TotalRoundsMapDeAztec                     string `db:"total_rounds_map_de_aztec"`
	TotalRoundsMapDeBank                      string `db:"total_rounds_map_de_bank"`
	TotalRoundsMapDeCbble                     string `db:"total_rounds_map_de_cbble"`
	TotalRoundsMapDeDust                      string `db:"total_rounds_map_de_dust"`
	TotalRoundsMapDeDust2                     string `db:"total_rounds_map_de_dust_2"`
	TotalRoundsMapDeInferno                   string `db:"total_rounds_map_de_inferno"`
	TotalRoundsMapDeLake                      string `db:"total_rounds_map_de_lake"`
	TotalRoundsMapDeNuke                      string `db:"total_rounds_map_de_nuke"`
	TotalRoundsMapDeSafehouse                 string `db:"total_rounds_map_de_safehouse"`
	TotalRoundsMapDeShorttrain                string `db:"total_rounds_map_de_shorttrain"`
	TotalRoundsMapDeStmarc                    string `db:"total_rounds_map_de_stmarc"`
	TotalRoundsMapDeSugarcane                 string `db:"total_rounds_map_de_sugarcane"`
	TotalRoundsMapDeTrain                     string `db:"total_rounds_map_de_train"`
	TotalRoundsMapDeVertigo                   string `db:"total_rounds_map_de_vertigo"`
	TotalRoundsPlayed                         string `db:"total_rounds_played"`
	TotalShotsAk47                            string `db:"total_shots_ak47"`
	TotalShotsAug                             string `db:"total_shots_aug"`
	TotalShotsAwp                             string `db:"total_shots_awp"`
	TotalShotsBizon                           string `db:"total_shots_bizon"`
	TotalShotsDeagle                          string `db:"total_shots_deagle"`
	TotalShotsElite                           string `db:"total_shots_elite"`
	TotalShotsFamas                           string `db:"total_shots_famas"`
	TotalShotsFired                           string `db:"total_shots_fired"`
	TotalShotsFiveseven                       string `db:"total_shots_fiveseven"`
	TotalShotsG3sg1                           string `db:"total_shots_g3sg1"`
	TotalShotsGalilar                         string `db:"total_shots_galilar"`
	TotalShotsGlock                           string `db:"total_shots_glock"`
	TotalShotsHit                             string `db:"total_shots_hit"`
	TotalShotsHkp2000                         string `db:"total_shots_hkp2000"`
	TotalShotsM249                            string `db:"total_shots_m249"`
	TotalShotsM4a1                            string `db:"total_shots_m4a1"`
	TotalShotsMac10                           string `db:"total_shots_mac10"`
	TotalShotsMag7                            string `db:"total_shots_mag7"`
	TotalShotsMp7                             string `db:"total_shots_mp7"`
	TotalShotsMp9                             string `db:"total_shots_mp9"`
	TotalShotsNegev                           string `db:"total_shots_negev"`
	TotalShotsNova                            string `db:"total_shots_nova"`
	TotalShotsP250                            string `db:"total_shots_p250"`
	TotalShotsP90                             string `db:"total_shots_p90"`
	TotalShotsSawedoff                        string `db:"total_shots_sawedoff"`
	TotalShotsScar20                          string `db:"total_shots_scar20"`
	TotalShotsSg556                           string `db:"total_shots_sg556"`
	TotalShotsSsg08                           string `db:"total_shots_ssg08"`
	TotalShotsTaser                           string `db:"total_shots_taser"`
	TotalShotsTec9                            string `db:"total_shots_tec9"`
	TotalShotsUmp45                           string `db:"total_shots_ump45"`
	TotalShotsXm1014                          string `db:"total_shots_xm1014"`
	TotalTRDefusedBombs                       string `db:"total_tr_defused_bombs"`
	TotalTRPlantedBombs                       string `db:"total_tr_planted_bombs"`
	TotalTimePlayed                           string `db:"total_time_played"`
	TotalTrbombMatchesWon                     string `db:"total_trbomb_matches_won"`
	TotalWeaponsDonated                       string `db:"total_weapons_donated"`
	TotalWins                                 string `db:"total_wins"`
	TotalWinsMapArBaggage                     string `db:"total_wins_map_ar_baggage"`
	TotalWinsMapArMonastery                   string `db:"total_wins_map_ar_monastery"`
	TotalWinsMapArShoots                      string `db:"total_wins_map_ar_shoots"`
	TotalWinsMapCsAssault                     string `db:"total_wins_map_cs_assault"`
	TotalWinsMapCsItaly                       string `db:"total_wins_map_cs_italy"`
	TotalWinsMapCsMilitia                     string `db:"total_wins_map_cs_militia"`
	TotalWinsMapCsOffice                      string `db:"total_wins_map_cs_office"`
	TotalWinsMapDeAztec                       string `db:"total_wins_map_de_aztec"`
	TotalWinsMapDeBank                        string `db:"total_wins_map_de_bank"`
	TotalWinsMapDeCbble                       string `db:"total_wins_map_de_cbble"`
	TotalWinsMapDeDust                        string `db:"total_wins_map_de_dust"`
	TotalWinsMapDeDust2                       string `db:"total_wins_map_de_dust_2"`
	TotalWinsMapDeHouse                       string `db:"total_wins_map_de_house"`
	TotalWinsMapDeInferno                     string `db:"total_wins_map_de_inferno"`
	TotalWinsMapDeLake                        string `db:"total_wins_map_de_lake"`
	TotalWinsMapDeNuke                        string `db:"total_wins_map_de_nuke"`
	TotalWinsMapDeSafehouse                   string `db:"total_wins_map_de_safehouse"`
	TotalWinsMapDeShorttrain                  string `db:"total_wins_map_de_shorttrain"`
	TotalWinsMapDeStmarc                      string `db:"total_wins_map_de_stmarc"`
	TotalWinsMapDeSugarcane                   string `db:"total_wins_map_de_sugarcane"`
	TotalWinsMapDeTrain                       string `db:"total_wins_map_de_train"`
	TotalWinsMapDeVertigo                     string `db:"total_wins_map_de_vertigo"`
	TotalWinsPistolround                      string `db:"total_wins_pistolround"`
}
