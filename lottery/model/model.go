package model

type PKTen struct {
	Periods    int
	FirstInfos []*FirstInfo
}

type FirstInfo struct {
	Number      int
	Miss        int
	MAXMiss     int
	AvarageMiss int
}

type ChipIn struct {
	FirstChips []int
	Times      int
	Principal  int
	Reward     float64
	Win        bool
	TimesToWin int
}

func (*PKTen) ChangeInfos(fierstPeriods int, numbers []int, pkt *PKTen) *PKTen {
	pkt.Periods = pkt.Periods + 1
	for _, i := range numbers {
		for _, k := range pkt.FirstInfos {
			if i == k.Number {
				k.Miss = 0
				k.AvarageMiss = (k.AvarageMiss*(pkt.Periods-fierstPeriods) + k.Miss) / (pkt.Periods - fierstPeriods + 1)
				if k.Miss <= k.MAXMiss {
					k.MAXMiss = k.MAXMiss
				} else {
					k.MAXMiss = k.Miss
				}
			} else {
				k.Miss = k.Miss + 1
				k.AvarageMiss = (k.AvarageMiss*(pkt.Periods-fierstPeriods) + k.Miss) / (pkt.Periods - fierstPeriods + 1)
				if k.Miss <= k.MAXMiss {
					k.MAXMiss = k.MAXMiss
				} else {
					k.MAXMiss = k.Miss
				}
			}
		}

	}
	return pkt
}
