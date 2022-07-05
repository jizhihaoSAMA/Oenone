package crud

var FieldMapper = map[string][]string{
	"index":   {"house_name", "house_size", "floor_info", "house_price", "house_loc_area", "house_loc_bc", "house_neighborhood", "is_full_rent", "support_short_term_rent", "has_lift", "has_single_toilet", "has_single_balcony", "image_amount"},
	"profile": {"house_name", "status", "house_loc_area", "house_loc_bc", "house_neighborhood", "created_time"},
	"admin":   {},
}
