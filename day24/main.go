package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- Part One ---")
	part1 := []string{
		"9996", "9985", "9974", "9963", "9952", "9941", "9896", "9885", "9874", "9863", "9852", "9841", "9796", "9785", "9774", "9763", "9752", "9741", "9696", "9685", "9674", "9663", "9652", "9641", "9596", "9585", "9574", "9563", "9552", "9541", "9496", "9485", "9474", "9463", "9452", "9441", "9396", "9385", "9374", "9363", "9352", "9341", "9296", "9285", "9274", "9263", "9252", "9241", "9196", "9185", "9174", "9163", "9152", "9141", "8996", "8985", "8974", "8963", "8952", "8941", "8896", "8885", "8874", "8863", "8852", "8841", "8796", "8785", "8774", "8763", "8752", "8741", "8696", "8685", "8674", "8663", "8652", "8641", "8596", "8585", "8574", "8563", "8552", "8541", "8496", "8485", "8474", "8463", "8452", "8441", "8396", "8385", "8374", "8363", "8352", "8341", "8296", "8285", "8274", "8263", "8252", "8241", "8196", "8185", "8174", "8163", "8152", "8141", "7996", "7985", "7974", "7963", "7952", "7941", "7896", "7885", "7874", "7863", "7852", "7841", "7796", "7785", "7774", "7763", "7752", "7741", "7696", "7685", "7674", "7663", "7652", "7641", "7596", "7585", "7574", "7563", "7552", "7541", "7496", "7485", "7474", "7463", "7452", "7441", "7396", "7385", "7374", "7363", "7352", "7341", "7296", "7285", "7274", "7263", "7252", "7241", "7196", "7185", "7174", "7163", "7152", "7141", "6996", "6985", "6974", "6963", "6952", "6941", "6896", "6885", "6874", "6863", "6852", "6841", "6796", "6785", "6774", "6763", "6752", "6741", "6696", "6685", "6674", "6663", "6652", "6641", "6596", "6585", "6574", "6563", "6552", "6541", "6496", "6485", "6474", "6463", "6452", "6441", "6396", "6385", "6374", "6363", "6352", "6341", "6296", "6285", "6274", "6263", "6252", "6241", "6196", "6185", "6174", "6163", "6152", "6141", "5996", "5985", "5974", "5963", "5952", "5941", "5896", "5885", "5874", "5863", "5852", "5841", "5796", "5785", "5774", "5763", "5752", "5741", "5696", "5685", "5674", "5663", "5652", "5641", "5596", "5585", "5574", "5563", "5552", "5541", "5496", "5485", "5474", "5463", "5452", "5441", "5396", "5385", "5374", "5363", "5352", "5341", "5296", "5285", "5274", "5263", "5252", "5241", "5196", "5185", "5174", "5163", "5152", "5141", "4996", "4985", "4974", "4963", "4952", "4941", "4896", "4885", "4874", "4863", "4852", "4841", "4796", "4785", "4774", "4763", "4752", "4741", "4696", "4685", "4674", "4663", "4652", "4641", "4596", "4585", "4574", "4563", "4552", "4541", "4496", "4485", "4474", "4463", "4452", "4441", "4396", "4385", "4374", "4363", "4352", "4341", "4296", "4285", "4274", "4263", "4252", "4241", "4196", "4185", "4174", "4163", "4152", "4141", "3996", "3985", "3974", "3963", "3952", "3941", "3896", "3885", "3874", "3863", "3852", "3841", "3796", "3785", "3774", "3763", "3752", "3741", "3696", "3685", "3674", "3663", "3652", "3641", "3596", "3585", "3574", "3563", "3552", "3541", "3496", "3485", "3474", "3463", "3452", "3441", "3396", "3385", "3374", "3363", "3352", "3341", "3296", "3285", "3274", "3263", "3252", "3241", "3196", "3185", "3174", "3163", "3152", "3141", "2996", "2985", "2974", "2963", "2952", "2941", "2896", "2885", "2874", "2863", "2852", "2841", "2796", "2785", "2774", "2763", "2752", "2741", "2696", "2685", "2674", "2663", "2652", "2641", "2596", "2585", "2574", "2563", "2552", "2541", "2496", "2485", "2474", "2463", "2452", "2441", "2396", "2385", "2374", "2363", "2352", "2341", "2296", "2285", "2274", "2263", "2252", "2241", "2196", "2185", "2174", "2163", "2152", "2141", "1996", "1985", "1974", "1963", "1952", "1941", "1896", "1885", "1874", "1863", "1852", "1841", "1796", "1785", "1774", "1763", "1752", "1741", "1696", "1685", "1674", "1663", "1652", "1641", "1596", "1585", "1574", "1563", "1552", "1541", "1496", "1485", "1474", "1463", "1452", "1441", "1396", "1385", "1374", "1363", "1352", "1341", "1296", "1285", "1274", "1263", "1252", "1241", "1196", "1185", "1174", "1163", "1152", "1141",
	}

	part2 := []string{
		"9994", "9983", "9972", "9961", "9894", "9883", "9872", "9861", "9794", "9783", "9772", "9761", "9694", "9683", "9672", "9661", "9594", "9583", "9572", "9561", "9494", "9483", "9472", "9461", "9394", "9383", "9372", "9361", "9294", "9283", "9272", "9261", "9194", "9183", "9172", "9161", "8994", "8983", "8972", "8961", "8894", "8883", "8872", "8861", "8794", "8783", "8772", "8761", "8694", "8683", "8672", "8661", "8594", "8583", "8572", "8561", "8494", "8483", "8472", "8461", "8394", "8383", "8372", "8361", "8294", "8283", "8272", "8261", "8194", "8183", "8172", "8161", "7994", "7983", "7972", "7961", "7894", "7883", "7872", "7861", "7794", "7783", "7772", "7761", "7694", "7683", "7672", "7661", "7594", "7583", "7572", "7561", "7494", "7483", "7472", "7461", "7394", "7383", "7372", "7361", "7294", "7283", "7272", "7261", "7194", "7183", "7172", "7161", "6994", "6983", "6972", "6961", "6894", "6883", "6872", "6861", "6794", "6783", "6772", "6761", "6694", "6683", "6672", "6661", "6594", "6583", "6572", "6561", "6494", "6483", "6472", "6461", "6394", "6383", "6372", "6361", "6294", "6283", "6272", "6261", "6194", "6183", "6172", "6161", "5994", "5983", "5972", "5961", "5894", "5883", "5872", "5861", "5794", "5783", "5772", "5761", "5694", "5683", "5672", "5661", "5594", "5583", "5572", "5561", "5494", "5483", "5472", "5461", "5394", "5383", "5372", "5361", "5294", "5283", "5272", "5261", "5194", "5183", "5172", "5161", "4994", "4983", "4972", "4961", "4894", "4883", "4872", "4861", "4794", "4783", "4772", "4761", "4694", "4683", "4672", "4661", "4594", "4583", "4572", "4561", "4494", "4483", "4472", "4461", "4394", "4383", "4372", "4361", "4294", "4283", "4272", "4261", "4194", "4183", "4172", "4161", "3994", "3983", "3972", "3961", "3894", "3883", "3872", "3861", "3794", "3783", "3772", "3761", "3694", "3683", "3672", "3661", "3594", "3583", "3572", "3561", "3494", "3483", "3472", "3461", "3394", "3383", "3372", "3361", "3294", "3283", "3272", "3261", "3194", "3183", "3172", "3161", "2994", "2983", "2972", "2961", "2894", "2883", "2872", "2861", "2794", "2783", "2772", "2761", "2694", "2683", "2672", "2661", "2594", "2583", "2572", "2561", "2494", "2483", "2472", "2461", "2394", "2383", "2372", "2361", "2294", "2283", "2272", "2261", "2194", "2183", "2172", "2161", "1994", "1983", "1972", "1961", "1894", "1883", "1872", "1861", "1794", "1783", "1772", "1761", "1694", "1683", "1672", "1661", "1594", "1583", "1572", "1561", "1494", "1483", "1472", "1461", "1394", "1383", "1372", "1361", "1294", "1283", "1272", "1261", "1194", "1183", "1172", "1161",
	}

	part3 := []string{
		"97", "86", "75", "64", "53", "42", "31",
	}

	i := 0
search:
	for _, p1 := range part1 {
		for _, p2 := range part2 {
			for _, p3 := range part3 {
				fmt.Println(i, len(part1)*len(part2)*len(part3))
				i++
				for sn := 9_999; sn >= 1_000; sn -= 10 {
					text := p1 + p2 + p3 + strconv.Itoa(sn)
					if strings.Contains(text, "0") {
						continue
					}
					input := arrayToInt(strings.Split(text, ""))

					as := []int{11, 11, 15, -11, 15, 15, 14, -7, 12, -6, -10, -15, -9, 0}
					bs := []int{1, 1, 1, 26, 1, 1, 1, 26, 1, 26, 26, 26, 26, 26}
					cs := []int{6, 12, 8, 7, 7, 12, 2, 15, 4, 5, 12, 11, 13, 7}

					// var ts []int
					// var zs []int

					z := 0
					for i, d := range input {
						t := z%26 + as[i]
						z /= bs[i]
						if i == 13 && t < 10 && z == 0 {
							fmt.Printf("%s%d\n", text[:13], t)
							break search
						} else {
							if t != d {
								z = 26*z + (d + cs[i])
							}
						}
						// ts = append(ts, t)
						// zs = append(zs, z)
					}

					// if zs[9] < 1_000_000 {
					// 	//fmt.Println(zs[9])
					// }

					// if z == 0 {
					// fmt.Printf("%q, ", text[4:8])
					// }
				}
			}
		}
	}

	i = 0
	for p1i := len(part1) - 1; p1i >= 0; p1i-- {
		p1 := part1[p1i]
		for p2i := len(part2) - 1; p2i >= 0; p2i-- {
			p2 := part2[p2i]
			for p3i := len(part3) - 1; p3i >= 0; p3i-- {
				p3 := part3[p3i]
				fmt.Println(i, len(part1)*len(part2)*len(part3))
				i++
				for sn := 1_111; sn <= 9_999; sn += 10 {
					text := p1 + p2 + p3 + strconv.Itoa(sn)
					if strings.Contains(text, "0") {
						continue
					}
					input := arrayToInt(strings.Split(text, ""))

					as := []int{11, 11, 15, -11, 15, 15, 14, -7, 12, -6, -10, -15, -9, 0}
					bs := []int{1, 1, 1, 26, 1, 1, 1, 26, 1, 26, 26, 26, 26, 26}
					cs := []int{6, 12, 8, 7, 7, 12, 2, 15, 4, 5, 12, 11, 13, 7}

					z := 0
					for i, d := range input {
						t := z%26 + as[i]
						z /= bs[i]
						if i == 13 && t < 10 && z == 0 {
							fmt.Printf("%s%d\n", text[:13], t)
							return
						} else {
							if t != d {
								z = 26*z + (d + cs[i])
							}
						}
					}
				}
			}
		}
	}
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func arrayToInt(input []string) (output []int) {
	output = make([]int, len(input))
	for i, text := range input {
		output[i] = toInt(text)
	}
	return output
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}