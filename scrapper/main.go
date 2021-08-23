package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main(){
	url := "https://www.bonappetit.com/api/bundles"
	method := "POST"

	payload := strings.NewReader(`{"url":"/bundles/59d6af2dcd8bb1736b3fdbde/revisions/26/containers/59d6af2dcd8bb1736b3fdbdc/items?itemFormat=model&page=1","contentType":"recipeChannel"}`)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Add("accept", "application/json;")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Origin", "https://www.bonappetit.com")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Referer", "https://www.bonappetit.com/recipes")
	req.Header.Add("Accept-Language", "es,en-US;q=0.9,en;q=0.8,es-AR;q=0.7,es-ES;q=0.6,es-419;q=0.5")
	req.Header.Add("Cookie", "verso_bucket=80; CN_geo_country_code=AR; CN_xid=8cdc37cd-7b9a-4b57-8f59-cf82047f20e2; CN_xid_refresh=8cdc37cd-7b9a-4b57-8f59-cf82047f20e2; usprivacy=1---; cneplayercaptions=showing; CN_ad_block=1; cneplayercount=49; CN_visits_m=1630465200202%26vn%3D1; CN_in_visit_m=true; AMCVS_F7093025512D2B690A490D44%40AdobeOrg=1; AMCV_F7093025512D2B690A490D44%40AdobeOrg=-408604571%7CMCIDTS%7C18844%7CMCMID%7C20681666231893707711335593025201259358%7CMCAAMLH-1628696730%7C4%7CMCAAMB-1628696730%7CRKhpRz8krg2tLO6pguXWp5olkAcUniQYPHaMWWgdJ3xzPWQmdj0y%7CMCOPTOUT-1628099130s%7CNONE%7CvVersion%7C4.6.0; s_vnum_m=1630465200699%26vn%3D1; sinvisit_m=true; _fbp=fb.1.1628091930860.1653950621; _scid=0666284d-2f28-4fb6-9eb3-20c5517569f8; sID=b6334712-74e5-4e9a-9995-7184e4a61039; CN_sp=61c2ed1f-46f4-47a8-85da-086c005fffc0; CN_su=ec038945-44ea-468a-95ee-29fc53ab7b08; _hjid=7e65da4b-9a75-401d-9cf3-a377713fee92; _hjFirstSeen=1; _hjIncludedInSessionSample=0; _hjAbsoluteSessionInProgress=1; _parsely_session={%22sid%22:1%2C%22surl%22:%22https://www.bonappetit.com/%22%2C%22sref%22:%22%22%2C%22sts%22:1628091931256%2C%22slts%22:0}; _parsely_visitor={%22id%22:%22pid=30544a8b50d47af6594e4c101546a0be%22%2C%22session_count%22:1%2C%22last_session_ts%22:1628091931256}; CN_segments=co.w2112; _sctr=1|1628046000000; _ga=GA1.2.289350965.1628091933; _gid=GA1.2.5557604.1628091933; bounceClientVisit2809v=N4IgNgDiBcIBYBcEQM4FIDMBBNAmAYnvgO6kB0ARgPYB2AhhBAKYICWCZAxlQLZEgAaEACcYIQSBRMA5jADaAXQC+QA; __qca=P0-1784427868-1628091931165; s_fid=550803D0DB09F635-036661065EC8757C; cn_4dsgcache=; __gads=ID=c7dae65a4700344d:T=1628091934:S=ALNI_MYNrdGtXy4NXfP2PocxOe12Siy3Eg; AMP_TOKEN=%24NOT_FOUND; _gat_UA-8293713-4=1; _dc_gtm_UA-8293713-4=1; aamconde=conde%3Dsv%3BCN%3D764354; aamoptsegs=aam%3D226821; aam_uuid=20891107355663593701352596615142739847; OptanonConsent=isIABGlobal=false&datestamp=Wed+Aug+04+2021+12%3A45%3A37+GMT-0300+(Argentina+Standard+Time)&version=6.19.0&hosts=&consentId=efb0f34f-a6dc-4258-8157-b8fac9344ac9&interactionCount=1&landingPath=NotLandingPage&groups=C0001%3A1%2CC0005%3A1%2CC0003%3A1%2CC0004%3A1%2CC0002%3A1&AwaitingReconsent=false&geolocation=AR%3BC&isGpcEnabled=0; OptanonAlertBoxClosed=2021-08-04T15:45:37.821Z; sailthru_pageviews=2; s_depth=2; s_ppn=www.bonappetit.com%2Frecipes; s_pct=index; sailthru_content=ed69ae8b2487ce57cfa27b7b4e3150a4a2cac7f36e29ac66d6bc26fee9ff5ef3; sailthru_visitor=8d9c1f40-357d-4726-ac17-e60a80a59777; pID=060dc489-2057-4ddf-80bf-15e0d54136fc; _derived_epik=dj0yJnU9X2FmVFUzOGJicmpIQ2U2ZXZMbC1LamFQaks0SzRGMlgmbj1GQ08wZFBoSGtfb3JtT2lkM29MQWNRJm09MSZ0PUFBQUFBR0VLdGlJJnJtPTEmcnQ9QUFBQUFHRUt0aUk; _pin_unauth=dWlkPU5XRTNORGhoWmpFdE5HSmxNeTAwWVRrekxXRmpPVFF0WW1Sak1XTTBaak0yWVdRMQ; _pbjs_userid_consent_data=3524755945110770; _lr_retry_request=true; _lr_env_src_ats=false; _pubcid_sharedid=01FC7EQYNS8ZTKD300KEHZTRQT; _pubcid=01FC7EQYNS8ZTKD300KEHZTRQT; timeSpent=1628091980592; s_nr=1628091980597-New; s_sq=conde-bonappetit%3D%2526c.%2526a.%2526activitymap.%2526page%253Dwww.bonappetit.com%25252Frecipes%2526link%253DMORE%252520RECIPES%2526region%253Dreact-app%2526pageIDType%253D1%2526.activitymap%2526.a%2526.c; CN_geo_country_code=AR; CN_xid=8cdc37cd-7b9a-4b57-8f59-cf82047f20e2; CN_xid_refresh=8cdc37cd-7b9a-4b57-8f59-cf82047f20e2; verso_bucket=80")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on client execution")
		fmt.Println(err)
		return
	}

	var bonappetitRecipes BonappetitScrap
	decErr := json.NewDecoder(res.Body).
		Decode(&bonappetitRecipes)

	if decErr != nil {
		fmt.Println("Error decoding")
		fmt.Println(decErr)
		return
	}

	//TODO: Validate in DB if the recipe (external) ID is already inserted and prevent duplicates
	fmt.Println(bonappetitRecipes.Recipes[0].Id)
	fmt.Println(bonappetitRecipes.Recipes[1].Id)
	fmt.Println(bonappetitRecipes.Recipes[2].Id)

	defer res.Body.Close()
}






























type BonappetitScrap struct {
	Recipes []BonappetitRecipe `json:"items"`
}
type BonappetitRecipe struct {
	Groups []IngredientGroup `json:"ingredientGroups"`
	PreparationGroups []struct {
		Hed   string `json:"hed"`
		Steps []struct {
			Description []interface{} `json:"description"`
		} `json:"steps"`
		MicroSteps []interface{} `json:"microSteps"`
	} `json:"preparationGroups"`
	Title string `json:"hed"`
	Id string `json:"id"`
	Asd string `json:"asd"`
}

type IngredientGroup struct {
	Title string `json:"hed"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name string `json:"name"`
}