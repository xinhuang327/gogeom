package geom

import (
	"math"
	"reflect"
)

func similar(a, b, e float64) bool {
	return math.Abs(a-b) < e
}

func pointSimilar(p1, p2 Point, e float64) bool {
	return similar(p1.X, p2.X, e) && similar(p1.Y, p2.Y, e)
}

func pointsSimilar(p1s, p2s []Point, e float64) bool {
	if len(p1s) != len(p2s) {
		return false
	}
	for i, n := 0, len(p1s); i < n; i++ {
		if !pointSimilar(p1s[i], p2s[i], e) {
			return false
		}
	}
	return true
}

func pointssSimilar(p1ss, p2ss [][]Point, e float64) bool {
	if len(p1ss) != len(p2ss) {
		return false
	}
	for i, n := 0, len(p1ss); i < n; i++ {
		if !pointsSimilar(p1ss[i], p2ss[i], e) {
			return false
		}
	}
	return true
}

func pointZSimilar(p1, p2 PointZ, e float64) bool {
	return similar(p1.X, p2.X, e) && similar(p1.Y, p2.Y, e) && similar(p1.Z, p2.Z, e)
}

func pointZsSimilar(p1s, p2s []PointZ, e float64) bool {
	if len(p1s) != len(p2s) {
		return false
	}
	for i, n := 0, len(p1s); i < n; i++ {
		if !pointZSimilar(p1s[i], p2s[i], e) {
			return false
		}
	}
	return true
}

func pointZssSimilar(p1ss, p2ss [][]PointZ, e float64) bool {
	if len(p1ss) != len(p2ss) {
		return false
	}
	for i, n := 0, len(p1ss); i < n; i++ {
		if !pointZsSimilar(p1ss[i], p2ss[i], e) {
			return false
		}
	}
	return true
}

func pointMSimilar(p1, p2 PointM, e float64) bool {
	return similar(p1.X, p2.X, e) && similar(p1.Y, p2.Y, e) && similar(p1.M, p2.M, e)
}

func pointMsSimilar(p1s, p2s []PointM, e float64) bool {
	if len(p1s) != len(p2s) {
		return false
	}
	for i, n := 0, len(p1s); i < n; i++ {
		if !pointMSimilar(p1s[i], p2s[i], e) {
			return false
		}
	}
	return true
}

func pointMssSimilar(p1ss, p2ss [][]PointM, e float64) bool {
	if len(p1ss) != len(p2ss) {
		return false
	}
	for i, n := 0, len(p1ss); i < n; i++ {
		if !pointMsSimilar(p1ss[i], p2ss[i], e) {
			return false
		}
	}
	return true
}

func pointZMSimilar(p1, p2 PointZM, e float64) bool {
	return similar(p1.X, p2.X, e) && similar(p1.Y, p2.Y, e) && similar(p1.Z, p2.Z, e) && similar(p1.M, p2.M, e)
}

func pointZMsSimilar(p1s, p2s []PointZM, e float64) bool {
	if len(p1s) != len(p2s) {
		return false
	}
	for i, n := 0, len(p1s); i < n; i++ {
		if !pointZMSimilar(p1s[i], p2s[i], e) {
			return false
		}
	}
	return true
}

func pointZMssSimilar(p1ss, p2ss [][]PointZM, e float64) bool {
	if len(p1ss) != len(p2ss) {
		return false
	}
	for i, n := 0, len(p1ss); i < n; i++ {
		if !pointZMsSimilar(p1ss[i], p2ss[i], e) {
			return false
		}
	}
	return true
}

func Similar(g1, g2 T, e float64) bool {
	if reflect.TypeOf(g1) != reflect.TypeOf(g2) {
		return false
	}
	switch g1.(type) {
	case Point:
		return pointSimilar(g1.(Point), g2.(Point), e)
	case PointZ:
		return pointZSimilar(g1.(PointZ), g2.(PointZ), e)
	case PointM:
		return pointMSimilar(g1.(PointM), g2.(PointM), e)
	case PointZM:
		return pointZMSimilar(g1.(PointZM), g2.(PointZM), e)
	case LineString:
		return pointsSimilar(g1.(LineString).Points, g2.(LineString).Points, e)
	case LineStringZ:
		return pointZsSimilar(g1.(LineStringZ).Points, g2.(LineStringZ).Points, e)
	case LineStringM:
		return pointMsSimilar(g1.(LineStringM).Points, g2.(LineStringM).Points, e)
	case LineStringZM:
		return pointZMsSimilar(g1.(LineStringZM).Points, g2.(LineStringZM).Points, e)
	case Polygon:
		return pointssSimilar(g1.(Polygon).Rings, g2.(Polygon).Rings, e)
	case PolygonZ:
		return pointZssSimilar(g1.(PolygonZ).Rings, g2.(PolygonZ).Rings, e)
	case PolygonM:
		return pointMssSimilar(g1.(PolygonM).Rings, g2.(PolygonM).Rings, e)
	case PolygonZM:
		return pointZMssSimilar(g1.(PolygonZM).Rings, g2.(PolygonZM).Rings, e)
	case MultiPoint:
		return pointsSimilar(g1.(MultiPoint).Points, g2.(MultiPoint).Points, e)
	case MultiPointZ:
		return pointZsSimilar(g1.(MultiPointZ).Points, g2.(MultiPointZ).Points, e)
	case MultiPointM:
		return pointMsSimilar(g1.(MultiPointM).Points, g2.(MultiPointM).Points, e)
	case MultiPointZM:
		return pointZMsSimilar(g1.(MultiPointZM).Points, g2.(MultiPointZM).Points, e)
	default:
		return false
	}
}
