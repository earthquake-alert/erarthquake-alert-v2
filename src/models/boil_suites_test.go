// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Areas", testAreas)
	t.Run("EarthquakeActivities", testEarthquakeActivities)
	t.Run("EarthquakeCounts", testEarthquakeCounts)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicenters)
	t.Run("EarthquakeInfos", testEarthquakeInfos)
	t.Run("EarthquakeReports", testEarthquakeReports)
	t.Run("EarthquakeUpdates", testEarthquakeUpdates)
	t.Run("Earthquakes", testEarthquakes)
	t.Run("IntensityStations", testIntensityStations)
	t.Run("JmaXmlEntries", testJmaXmlEntries)
	t.Run("Prefectures", testPrefectures)
	t.Run("TsunamiConnects", testTsunamiConnects)
	t.Run("TsunamiInfos", testTsunamiInfos)
	t.Run("TwitterThreads", testTwitterThreads)
}

func TestDelete(t *testing.T) {
	t.Run("Areas", testAreasDelete)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesDelete)
	t.Run("EarthquakeCounts", testEarthquakeCountsDelete)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersDelete)
	t.Run("EarthquakeInfos", testEarthquakeInfosDelete)
	t.Run("EarthquakeReports", testEarthquakeReportsDelete)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesDelete)
	t.Run("Earthquakes", testEarthquakesDelete)
	t.Run("IntensityStations", testIntensityStationsDelete)
	t.Run("JmaXmlEntries", testJmaXmlEntriesDelete)
	t.Run("Prefectures", testPrefecturesDelete)
	t.Run("TsunamiConnects", testTsunamiConnectsDelete)
	t.Run("TsunamiInfos", testTsunamiInfosDelete)
	t.Run("TwitterThreads", testTwitterThreadsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Areas", testAreasQueryDeleteAll)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesQueryDeleteAll)
	t.Run("EarthquakeCounts", testEarthquakeCountsQueryDeleteAll)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersQueryDeleteAll)
	t.Run("EarthquakeInfos", testEarthquakeInfosQueryDeleteAll)
	t.Run("EarthquakeReports", testEarthquakeReportsQueryDeleteAll)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesQueryDeleteAll)
	t.Run("Earthquakes", testEarthquakesQueryDeleteAll)
	t.Run("IntensityStations", testIntensityStationsQueryDeleteAll)
	t.Run("JmaXmlEntries", testJmaXmlEntriesQueryDeleteAll)
	t.Run("Prefectures", testPrefecturesQueryDeleteAll)
	t.Run("TsunamiConnects", testTsunamiConnectsQueryDeleteAll)
	t.Run("TsunamiInfos", testTsunamiInfosQueryDeleteAll)
	t.Run("TwitterThreads", testTwitterThreadsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Areas", testAreasSliceDeleteAll)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesSliceDeleteAll)
	t.Run("EarthquakeCounts", testEarthquakeCountsSliceDeleteAll)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersSliceDeleteAll)
	t.Run("EarthquakeInfos", testEarthquakeInfosSliceDeleteAll)
	t.Run("EarthquakeReports", testEarthquakeReportsSliceDeleteAll)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesSliceDeleteAll)
	t.Run("Earthquakes", testEarthquakesSliceDeleteAll)
	t.Run("IntensityStations", testIntensityStationsSliceDeleteAll)
	t.Run("JmaXmlEntries", testJmaXmlEntriesSliceDeleteAll)
	t.Run("Prefectures", testPrefecturesSliceDeleteAll)
	t.Run("TsunamiConnects", testTsunamiConnectsSliceDeleteAll)
	t.Run("TsunamiInfos", testTsunamiInfosSliceDeleteAll)
	t.Run("TwitterThreads", testTwitterThreadsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Areas", testAreasExists)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesExists)
	t.Run("EarthquakeCounts", testEarthquakeCountsExists)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersExists)
	t.Run("EarthquakeInfos", testEarthquakeInfosExists)
	t.Run("EarthquakeReports", testEarthquakeReportsExists)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesExists)
	t.Run("Earthquakes", testEarthquakesExists)
	t.Run("IntensityStations", testIntensityStationsExists)
	t.Run("JmaXmlEntries", testJmaXmlEntriesExists)
	t.Run("Prefectures", testPrefecturesExists)
	t.Run("TsunamiConnects", testTsunamiConnectsExists)
	t.Run("TsunamiInfos", testTsunamiInfosExists)
	t.Run("TwitterThreads", testTwitterThreadsExists)
}

func TestFind(t *testing.T) {
	t.Run("Areas", testAreasFind)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesFind)
	t.Run("EarthquakeCounts", testEarthquakeCountsFind)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersFind)
	t.Run("EarthquakeInfos", testEarthquakeInfosFind)
	t.Run("EarthquakeReports", testEarthquakeReportsFind)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesFind)
	t.Run("Earthquakes", testEarthquakesFind)
	t.Run("IntensityStations", testIntensityStationsFind)
	t.Run("JmaXmlEntries", testJmaXmlEntriesFind)
	t.Run("Prefectures", testPrefecturesFind)
	t.Run("TsunamiConnects", testTsunamiConnectsFind)
	t.Run("TsunamiInfos", testTsunamiInfosFind)
	t.Run("TwitterThreads", testTwitterThreadsFind)
}

func TestBind(t *testing.T) {
	t.Run("Areas", testAreasBind)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesBind)
	t.Run("EarthquakeCounts", testEarthquakeCountsBind)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersBind)
	t.Run("EarthquakeInfos", testEarthquakeInfosBind)
	t.Run("EarthquakeReports", testEarthquakeReportsBind)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesBind)
	t.Run("Earthquakes", testEarthquakesBind)
	t.Run("IntensityStations", testIntensityStationsBind)
	t.Run("JmaXmlEntries", testJmaXmlEntriesBind)
	t.Run("Prefectures", testPrefecturesBind)
	t.Run("TsunamiConnects", testTsunamiConnectsBind)
	t.Run("TsunamiInfos", testTsunamiInfosBind)
	t.Run("TwitterThreads", testTwitterThreadsBind)
}

func TestOne(t *testing.T) {
	t.Run("Areas", testAreasOne)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesOne)
	t.Run("EarthquakeCounts", testEarthquakeCountsOne)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersOne)
	t.Run("EarthquakeInfos", testEarthquakeInfosOne)
	t.Run("EarthquakeReports", testEarthquakeReportsOne)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesOne)
	t.Run("Earthquakes", testEarthquakesOne)
	t.Run("IntensityStations", testIntensityStationsOne)
	t.Run("JmaXmlEntries", testJmaXmlEntriesOne)
	t.Run("Prefectures", testPrefecturesOne)
	t.Run("TsunamiConnects", testTsunamiConnectsOne)
	t.Run("TsunamiInfos", testTsunamiInfosOne)
	t.Run("TwitterThreads", testTwitterThreadsOne)
}

func TestAll(t *testing.T) {
	t.Run("Areas", testAreasAll)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesAll)
	t.Run("EarthquakeCounts", testEarthquakeCountsAll)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersAll)
	t.Run("EarthquakeInfos", testEarthquakeInfosAll)
	t.Run("EarthquakeReports", testEarthquakeReportsAll)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesAll)
	t.Run("Earthquakes", testEarthquakesAll)
	t.Run("IntensityStations", testIntensityStationsAll)
	t.Run("JmaXmlEntries", testJmaXmlEntriesAll)
	t.Run("Prefectures", testPrefecturesAll)
	t.Run("TsunamiConnects", testTsunamiConnectsAll)
	t.Run("TsunamiInfos", testTsunamiInfosAll)
	t.Run("TwitterThreads", testTwitterThreadsAll)
}

func TestCount(t *testing.T) {
	t.Run("Areas", testAreasCount)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesCount)
	t.Run("EarthquakeCounts", testEarthquakeCountsCount)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersCount)
	t.Run("EarthquakeInfos", testEarthquakeInfosCount)
	t.Run("EarthquakeReports", testEarthquakeReportsCount)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesCount)
	t.Run("Earthquakes", testEarthquakesCount)
	t.Run("IntensityStations", testIntensityStationsCount)
	t.Run("JmaXmlEntries", testJmaXmlEntriesCount)
	t.Run("Prefectures", testPrefecturesCount)
	t.Run("TsunamiConnects", testTsunamiConnectsCount)
	t.Run("TsunamiInfos", testTsunamiInfosCount)
	t.Run("TwitterThreads", testTwitterThreadsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Areas", testAreasHooks)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesHooks)
	t.Run("EarthquakeCounts", testEarthquakeCountsHooks)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersHooks)
	t.Run("EarthquakeInfos", testEarthquakeInfosHooks)
	t.Run("EarthquakeReports", testEarthquakeReportsHooks)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesHooks)
	t.Run("Earthquakes", testEarthquakesHooks)
	t.Run("IntensityStations", testIntensityStationsHooks)
	t.Run("JmaXmlEntries", testJmaXmlEntriesHooks)
	t.Run("Prefectures", testPrefecturesHooks)
	t.Run("TsunamiConnects", testTsunamiConnectsHooks)
	t.Run("TsunamiInfos", testTsunamiInfosHooks)
	t.Run("TwitterThreads", testTwitterThreadsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Areas", testAreasInsert)
	t.Run("Areas", testAreasInsertWhitelist)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesInsert)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesInsertWhitelist)
	t.Run("EarthquakeCounts", testEarthquakeCountsInsert)
	t.Run("EarthquakeCounts", testEarthquakeCountsInsertWhitelist)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersInsert)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersInsertWhitelist)
	t.Run("EarthquakeInfos", testEarthquakeInfosInsert)
	t.Run("EarthquakeInfos", testEarthquakeInfosInsertWhitelist)
	t.Run("EarthquakeReports", testEarthquakeReportsInsert)
	t.Run("EarthquakeReports", testEarthquakeReportsInsertWhitelist)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesInsert)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesInsertWhitelist)
	t.Run("Earthquakes", testEarthquakesInsert)
	t.Run("Earthquakes", testEarthquakesInsertWhitelist)
	t.Run("IntensityStations", testIntensityStationsInsert)
	t.Run("IntensityStations", testIntensityStationsInsertWhitelist)
	t.Run("JmaXmlEntries", testJmaXmlEntriesInsert)
	t.Run("JmaXmlEntries", testJmaXmlEntriesInsertWhitelist)
	t.Run("Prefectures", testPrefecturesInsert)
	t.Run("Prefectures", testPrefecturesInsertWhitelist)
	t.Run("TsunamiConnects", testTsunamiConnectsInsert)
	t.Run("TsunamiConnects", testTsunamiConnectsInsertWhitelist)
	t.Run("TsunamiInfos", testTsunamiInfosInsert)
	t.Run("TsunamiInfos", testTsunamiInfosInsertWhitelist)
	t.Run("TwitterThreads", testTwitterThreadsInsert)
	t.Run("TwitterThreads", testTwitterThreadsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Areas", testAreasReload)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesReload)
	t.Run("EarthquakeCounts", testEarthquakeCountsReload)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersReload)
	t.Run("EarthquakeInfos", testEarthquakeInfosReload)
	t.Run("EarthquakeReports", testEarthquakeReportsReload)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesReload)
	t.Run("Earthquakes", testEarthquakesReload)
	t.Run("IntensityStations", testIntensityStationsReload)
	t.Run("JmaXmlEntries", testJmaXmlEntriesReload)
	t.Run("Prefectures", testPrefecturesReload)
	t.Run("TsunamiConnects", testTsunamiConnectsReload)
	t.Run("TsunamiInfos", testTsunamiInfosReload)
	t.Run("TwitterThreads", testTwitterThreadsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Areas", testAreasReloadAll)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesReloadAll)
	t.Run("EarthquakeCounts", testEarthquakeCountsReloadAll)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersReloadAll)
	t.Run("EarthquakeInfos", testEarthquakeInfosReloadAll)
	t.Run("EarthquakeReports", testEarthquakeReportsReloadAll)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesReloadAll)
	t.Run("Earthquakes", testEarthquakesReloadAll)
	t.Run("IntensityStations", testIntensityStationsReloadAll)
	t.Run("JmaXmlEntries", testJmaXmlEntriesReloadAll)
	t.Run("Prefectures", testPrefecturesReloadAll)
	t.Run("TsunamiConnects", testTsunamiConnectsReloadAll)
	t.Run("TsunamiInfos", testTsunamiInfosReloadAll)
	t.Run("TwitterThreads", testTwitterThreadsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Areas", testAreasSelect)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesSelect)
	t.Run("EarthquakeCounts", testEarthquakeCountsSelect)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersSelect)
	t.Run("EarthquakeInfos", testEarthquakeInfosSelect)
	t.Run("EarthquakeReports", testEarthquakeReportsSelect)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesSelect)
	t.Run("Earthquakes", testEarthquakesSelect)
	t.Run("IntensityStations", testIntensityStationsSelect)
	t.Run("JmaXmlEntries", testJmaXmlEntriesSelect)
	t.Run("Prefectures", testPrefecturesSelect)
	t.Run("TsunamiConnects", testTsunamiConnectsSelect)
	t.Run("TsunamiInfos", testTsunamiInfosSelect)
	t.Run("TwitterThreads", testTwitterThreadsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Areas", testAreasUpdate)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesUpdate)
	t.Run("EarthquakeCounts", testEarthquakeCountsUpdate)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersUpdate)
	t.Run("EarthquakeInfos", testEarthquakeInfosUpdate)
	t.Run("EarthquakeReports", testEarthquakeReportsUpdate)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesUpdate)
	t.Run("Earthquakes", testEarthquakesUpdate)
	t.Run("IntensityStations", testIntensityStationsUpdate)
	t.Run("JmaXmlEntries", testJmaXmlEntriesUpdate)
	t.Run("Prefectures", testPrefecturesUpdate)
	t.Run("TsunamiConnects", testTsunamiConnectsUpdate)
	t.Run("TsunamiInfos", testTsunamiInfosUpdate)
	t.Run("TwitterThreads", testTwitterThreadsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Areas", testAreasSliceUpdateAll)
	t.Run("EarthquakeActivities", testEarthquakeActivitiesSliceUpdateAll)
	t.Run("EarthquakeCounts", testEarthquakeCountsSliceUpdateAll)
	t.Run("EarthquakeEpicenters", testEarthquakeEpicentersSliceUpdateAll)
	t.Run("EarthquakeInfos", testEarthquakeInfosSliceUpdateAll)
	t.Run("EarthquakeReports", testEarthquakeReportsSliceUpdateAll)
	t.Run("EarthquakeUpdates", testEarthquakeUpdatesSliceUpdateAll)
	t.Run("Earthquakes", testEarthquakesSliceUpdateAll)
	t.Run("IntensityStations", testIntensityStationsSliceUpdateAll)
	t.Run("JmaXmlEntries", testJmaXmlEntriesSliceUpdateAll)
	t.Run("Prefectures", testPrefecturesSliceUpdateAll)
	t.Run("TsunamiConnects", testTsunamiConnectsSliceUpdateAll)
	t.Run("TsunamiInfos", testTsunamiInfosSliceUpdateAll)
	t.Run("TwitterThreads", testTwitterThreadsSliceUpdateAll)
}
