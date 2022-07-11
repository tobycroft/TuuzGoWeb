package Redis

import (
	"github.com/go-redis/redis/v8"
	"main.go/config/app_conf"
)

func Geo_add(key string, longitude, latitude float64, geo_name string) error {
	var geo redis.GeoLocation
	geo.Longitude = longitude
	geo.Latitude = latitude
	geo.Name = geo_name
	return goredis.GeoAdd(goredis_ctx, app_conf.Project+":"+key, &geo).Err()
}

func Geo_get_pos(key string, geo_name string) (*redis.GeoPos, error) {
	poss, err := goredis.GeoPos(goredis_ctx, key, geo_name).Result()
	if err != nil {
		return nil, err
	}
	return poss[0], err
}

func Geo_get_poss(key string, geo_names ...string) ([]*redis.GeoPos, error) {
	return goredis.GeoPos(goredis_ctx, app_conf.Project+":"+key, geo_names...).Result()
}

func Geo_distance(key string, geo_name1, geo_name2 string, in_m_km_ft_mi string) (float64, error) {
	return goredis.GeoDist(goredis_ctx, app_conf.Project+":"+key, geo_name1, geo_name2, in_m_km_ft_mi).Result()
}

func Geo_range(key string, lon, lat float64, area float64, in_m_km_ft_mi string, count int) ([]redis.GeoLocation, error) {
	var geo redis.GeoRadiusQuery
	geo.Count = count
	geo.Sort = "ASC"
	geo.Radius = area
	geo.WithDist = true
	geo.WithCoord = true
	geo.Unit = in_m_km_ft_mi
	return goredis.GeoRadius(goredis_ctx, app_conf.Project+":"+key, lon, lat, &geo).Result()
}

func Geo_range_byGeoName(key string, geo_name string, area float64, in_m_km_ft_mi string, count int) ([]redis.GeoLocation, error) {
	var geo redis.GeoRadiusQuery
	geo.Radius = area
	geo.Unit = in_m_km_ft_mi
	geo.Count = count
	return goredis.GeoRadiusByMember(goredis_ctx, app_conf.Project+":"+key, geo_name, &geo).Result()
}

func Geo_search_geoName(key string, lon, lat float64, area float64, in_m_km_ft_mi string, count int) ([]string, error) {
	var geo redis.GeoSearchQuery
	geo.Longitude = lon
	geo.Latitude = lat
	geo.Radius = area
	geo.RadiusUnit = in_m_km_ft_mi
	geo.Sort = "ASC"
	geo.Count = count
	return goredis.GeoSearch(goredis_ctx, app_conf.Project+":"+key, &geo).Result()
}

func Geo_search_location_byName(key string, geo_name string, area float64, in_m_km_ft_mi string, count int) ([]redis.GeoLocation, error) {
	var geo redis.GeoSearchLocationQuery
	geo.WithCoord = true
	geo.WithDist = true

	geo.Member = geo_name
	geo.Radius = area
	geo.RadiusUnit = in_m_km_ft_mi

	geo.Count = count
	geo.Sort = "ASC"
	return goredis.GeoSearchLocation(goredis_ctx, app_conf.Project+":"+key, &geo).Result()
}

func Geo_search_location_byLonLat(key string, lon, lat, area float64, in_m_km_ft_mi string, count int) ([]redis.GeoLocation, error) {
	var geo redis.GeoSearchLocationQuery
	geo.WithCoord = true
	geo.WithDist = true

	geo.Longitude = lon
	geo.Latitude = lat
	geo.Radius = area
	geo.RadiusUnit = in_m_km_ft_mi

	geo.Count = count
	geo.Sort = "ASC"

	return goredis.GeoSearchLocation(goredis_ctx, app_conf.Project+":"+key, &geo).Result()
}

func Geo_search_location_byLonLat_toStore(key, to_key string, lon, lat, area float64, in_m_km_ft_mi string, count int) error {
	var geo redis.GeoSearchStoreQuery

	geo.Longitude = lon
	geo.Latitude = lat
	geo.Radius = area
	geo.RadiusUnit = in_m_km_ft_mi

	geo.Count = count
	geo.Sort = "ASC"

	return goredis.GeoSearchStore(goredis_ctx, app_conf.Project+":"+key, app_conf.Project+":"+to_key, &geo).Err()
}

func Geo_search_location_byName_toStore(key, to_key string, geo_name string, area float64, in_m_km_ft_mi string, count int) error {
	var geo redis.GeoSearchStoreQuery

	geo.Member = geo_name
	geo.Radius = area
	geo.RadiusUnit = in_m_km_ft_mi

	geo.Count = count
	geo.Sort = "ASC"

	return goredis.GeoSearchStore(goredis_ctx, app_conf.Project+":"+key, app_conf.Project+":"+to_key, &geo).Err()
}
