package main

type PresetLocations []string

var PRESETS = map[string]PresetLocations{
  "austria":PresetLocations{"austria", "%C3%B6sterreich", "vienna", "wien", "linz", "salzburg", "graz", "innsbruck", "klagenfurt", "wels", "dornbirn"},
  "finland":PresetLocations{"finland", "suomi", "helsinki", "tampere", "oulu", "espoo", "vantaa", "turku"},
  "sweden":PresetLocations{"sweden", "sverige", "stockholm", "malm%C3%B6", "uppsala", "g%C3%B6teborg", "gothenburg"},
  "norway":PresetLocations{"norway", "norge", "oslo", "bergen"},
  "germany":PresetLocations{"germany", "deutschland", "berlin", "frankfurt", "munich", "m%C3%BCnchen", "hamburg", "cologne", "k%C3%B6ln"},
  "netherlands":PresetLocations{"netherlands", "nederland", "amsterdam", "rotterdam", "hague", "utrecht", "holland", "delft"},
  "ukraine":PresetLocations{"ukraine", "kiev", "kyiv", "kharkiv", "dnipro", "odesa", "donetsk", "zaporizhia"},
  "japan":PresetLocations{"japan", "tokyo", "yokohama", "osaka", "nagoya", "sapporo", "kobe", "kyoto", "fukuoka", "kawasaki", "saitama", "hiroshima", "sendai"},
  "russia":PresetLocations{"russia", "moscow", "saint%2Bpetersburg", "novosibirsk", "yekaterinburg", "nizhny%2Bnovgorod", "samara", "omsk", "kazan", "chelyabinsk", "rostov-on-don", "ufa", "volgograd"},
  "estonia":PresetLocations{"estonia", "eesti", "tallinn", "tartu", "narva", "p%C3%A4rnu"},
  "denmark":PresetLocations{"denmark", "danmark", "copenhagen","aarhus","odense","aalborg"},
  "portugal":PresetLocations{"portugal", "lisbon", "lisboa","braga","porto","aveiro","coimbra","funchal", "madeira"},
  "france":PresetLocations{"france","paris","marseille","lyon","toulouse","nice","nantes","strasbourg","montpellier","bordeaux","lille","rennes","reims"},
  "spain":PresetLocations{"spain","espa%C3%B1a","madrid","barcelona","valencia","seville","sevilla","zaragoza","malaga","murcia","palma","bilbao","alicante","cordoba"},
  "italy":PresetLocations{"italy","italia","rome","roma","milan","naples","napoli","turin","torino","palermo","genoa","genova","bologna","florence","firenze","bari","catania","venice","verona"},
  "uk": PresetLocations{"uk","london","birmingham","leeds","glasgow","sheffield","bradford","manchester","edinburgh","liverpool","bristol","cardiff","belfast","leicester","wakefield","coventry","nottingham","newcastle"},
  "croatia": PresetLocations{"croatia","hrvatska","zagreb","split","rijeka","osijek","zadar","pula"},
  "worldwide": PresetLocations{},
  "china": PresetLocations{"china", "%E4%B8%AD%E5%9B%BD", "guangzhou", "shanghai", "beijing", "hangzhou", "hong%2Bkong"},
  "india": PresetLocations{"india", "mumbai", "delhi", "bangalore", "hyderabad", "ahmedabad", "chennai", "kolkata", "jaipur"},
  "israel": PresetLocations{"israel", "tel%2Baviv", "jerusalem", "beer%2Bsheva", "beersheva", "netanya", "ramat%2Bgan", "haifa", "herzliya", "rishon"},
  "indonesia": PresetLocations{"indonesia", "jakarta", "surabaya", "bandung", "medan", "bekasi", "semarang", "tangerang", "depok", "makassar", "palembang"},
  "pakistan": PresetLocations{"pakistan", "karachi", "lahore", "faisalabad", "rawalpindi", "peshawar", "islamabad"},
  "brazil": PresetLocations{"brazil", "brasil", "s%C3%A3o%20paulo", "bras%C3%ADlia", "salvador", "fortaleza", "belo%20horizonte", "manaus", "curitiba", "recife", "porto%20alegre"},
  "nigeria": PresetLocations{"nigeria", "lagos", "kano", "ibadan", "benin%20city", "port%20harcourt", "jos", "ilorin"},
  "bangladesh": PresetLocations{"bangladesh", "dhaka", "chittagong", "khulna", "rajshahi", "barisal", "sylhet", "rangpur", "comilla", "gazipur"},
  "mexico": PresetLocations{"mexico", "mexico%20city", "guadalajara", "puebla", "tijuana", "mexicali"},
  "philippines": PresetLocations{"philippines", "pilipinas", "quezon", "manila", "davao", "caloocan", "cebu", "zamboanga", "bohol", "pasig", "bacolod", "makati", "baguio"},
  "luxembourg": PresetLocations{"luxembourg", "esch-sur-alzette", "differdange", "dudelange", "ettelbruck", "diekirch", "wiltz", "echternach", "rumelange", "grevenmacher", "bertrange", "mamer", "capellen", "strassen", "diekirch"},
  "egypt": PresetLocations{"egypt", "cairo", "alexandria", "giza", "port%2Bsaid", "suez", "luxor", "el%2Bmahalla", "asyut", "al%2Bmansurah", "tanda"},
  "ethiopia": PresetLocations{"ethiopia", "addis%2Bababa", "gondar", "adama", "hawassa", "bahir%2Bdar"},
  "vietnam": PresetLocations{"vietnam", "viet%2Bnam", "ho%2Bchi%2Bminh", "hanoi", "ha%2Bnoi", "hai%2Bphong", "da%2Bnang", "can%2Btho", "bien%2Bhoa", "nha%Btrang", "vinh"},
  "iran": PresetLocations{"iran", "tehran", "mashhad", "isfahan", "esfahan", "karaj", "shiraz", "tabriz", "qom", "ahvaz", "ahwaz", "kermanshah", "urmia", "rasht", "kerman"},
  "congo": PresetLocations{"congo", "drc", "kinshasa", "lubumbashi", "bukavu", "kananga"},
  "turkey": PresetLocations{"turkey", "turkiye", "istanbul", "ankara", "izmir", "bursa", "adana", "gaziantep", "konya", "antalya", "kayseri", "mersin", "eskisehir", "samsun", "denizli", "malatya"},
  "thailand": PresetLocations{"thailand", "bangkok", "nonthaburi", "nakhon", "phuket", "pattaya", "chiang%2Bmai"},
  "south africa": PresetLocations{"south%2Bafrica", "south%2Bafrica", "johannesburg", "cape%2Btown", "rsa", "durban", "port%2Belizabeth", "pretoria", "nelspruit"},
  "myanmar": PresetLocations{"myanmar", "burma", "yangon", "rangoon", "mandalay", "nay%2Bpyi%2Btaw", "taunggyi", "bago", "mawlamyine"},
  "tanzania": PresetLocations{"tanzania", "dar%2Bes%2Bsalaam", "mwanza", "arusha", "dodoma", "mbeya", "morogoro", "tanga", "kilimanjaro"},
  "south korea": PresetLocations{"south%2Bkorea", "ROK", "korea", "seoul", "busan", "incheon", "daegu", "daejeon", "gwangju", "%EB%8C%80%ED%95%9C%EB%AF%BC%EA%B5%AD", "%EC%84%9C%EC%9A%B8", "%EC%84%9C%EC%9A%B8%EC%8B%9C"},
  "colombia":PresetLocations{"colombia","bogota","medellin","cali","barranquilla","cartagena","cucuta","bucaramanga","ibague","soledad","pereira","santa%2Bmarta"},
  "kenya":PresetLocations{"kenya","nairobi","mombasa","kisumu","nakuru","eldoret","kisii"},
  "argentina":PresetLocations{"argentina","buenos%2Baires","cordoba","rosario","mendoza","la%2Bplata","tucuman","mar%2Bdel%2Bplata","salta","resistencia"},
  "algeria":PresetLocations{"algeria","algiers","oran","constantine","annaba","blida","batna","djelfa","setif","sidi%2Bbel%2Babbes","biskra"},
  "sudan":PresetLocations{"sudan","khartoum","omdurman"},
  "poland":PresetLocations{"poland","polska","warsaw","krakow","lodz","wroclaw","poznan","gdansk","szczecin","bydgoszcz","lublin","katowice","bialystok"},
  "canada":PresetLocations{"canada","ottawa","edmonton","winnipeg","vancouver","toronto","quebec","montreal","mississauga","calgary"},
  "australia":PresetLocations{"australia","sydney","melbourne","brisbane","perth","adelaide","canberra","hobart"},
  "belgium":PresetLocations{"belgium","antwerp","ghent","charleroi","liege","brussels","belgique"},
  "greece":PresetLocations{"greece","%CE%95%CE%BB%CE%BB%CE%AC%CE%B4%CE%B1","athens","thessaloniki","patras","heraklion","larissa","volos","rhodes","ioannina","chania","crete"}}

func Preset(name string) []string {
  return PRESETS[name]
}
