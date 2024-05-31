package countries

type Country struct {
	Name   string
	Alpha2 string
	Alpha3 string
}

func GetByAlpha3(alpha3 string) Country {
	if alpha2, ok := isoCodeMap[alpha3]; ok {
		return Country{
			Name:   names[alpha3],
			Alpha2: alpha2,
			Alpha3: alpha3,
		}
	}

	return Country{}
}

func GetByAlpha2(alpha2 string) Country {
	for alpha3, a2 := range isoCodeMap {
		if a2 == alpha2 {
			return Country{
				Name:   names[alpha3],
				Alpha2: alpha2,
				Alpha3: alpha3,
			}
		}
	}

	return Country{}
}

func GetByName(name string) Country {
	for alpha3, n := range names {
		if n == name {
			return Country{
				Name:   name,
				Alpha2: isoCodeMap[alpha3],
				Alpha3: alpha3,
			}
		}
	}

	return Country{}
}
