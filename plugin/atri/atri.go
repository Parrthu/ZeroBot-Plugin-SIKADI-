/*
Package atri 本文件基于 https://github.com/Kyomotoi/ATRI
为 Golang 移植版，语料、素材均来自上述项目
本项目遵守 AGPL v3 协议进行开源
*/
package atri

import (
	"math/rand"
	"time"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	control "github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/process"
)

const (
	// 服务名
	servicename = "atri"
	// ATRI 表情的 codechina 镜像
	res = "https://gitcode.net/u011570312/zbpdata/-/raw/main/Atri/"
)

func init() { // 插件主体
	engine := control.Register(servicename, &control.Options{
		DisableOnDefault: false,
		Help: "本插件基于 ATRI ，为 Golang 移植版，由parrthu修改nickname为斯卡蒂\n" +
			"- 斯卡蒂醒醒\n- 斯卡蒂睡吧\n- 萝卜子\n- 喜欢 | 爱你 | 爱 | suki | daisuki | すき | 好き | 贴贴 | 老婆 | 亲一个 | mua\n" +
			"- 草你妈 | 操你妈 | 脑瘫 | 废柴 | fw | 废物 | 战斗 | 爬 | 爪巴 | sb | SB | 傻B\n- 早安 | 早哇 | 早上好 | ohayo | 哦哈哟 | お早う | 早好 | 早 | 早早早\n" +
			"- 中午好 | 午安 | 午好\n- 晚安 | oyasuminasai | おやすみなさい | 晚好 | 晚上好\n- 高性能 | 太棒了 | すごい | sugoi | 斯国一 | よかった\n" +
			"- 没事 | 没关系 | 大丈夫 | 还好 | 不要紧 | 没出大问题 | 没伤到哪\n- 好吗 | 是吗 | 行不行 | 能不能 | 可不可以\n- 啊这\n- 我好了\n- ？ | ? | ¿\n" +
			"- 离谱\n- 答应我",
		OnEnable: func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(message.Text("嗯呜呜……夏生先生……？"))
		},
		OnDisable: func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(message.Text("Zzz……Zzz……"))
		},
	})
	engine.OnFullMatch("蒂蒂", isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			switch rand.Intn(2) {
			case 0:
				ctx.SendChain(randText("蒂蒂是对机器人的蔑称！", "是斯卡蒂......蒂蒂可是对机器人的蔑称"))
			case 1:
				ctx.SendChain(randRecord("RocketPunch.amr"))
			}
		})
	engine.OnFullMatchGroup([]string{"喜欢", "爱你", "爱", "suki", "daisuki", "すき", "好き", "贴贴", "老婆", "亲一个", "mua"}, isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(randImage("SUKI.jpg", "SUKI1.jpg", "SUKI2.png"))
		})
	engine.OnKeywordGroup([]string{"草你妈", "操你妈", "脑瘫", "废柴", "fw", "five", "废物", "战斗", "爬", "爪巴", "sb", "SB", "傻B"}, isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(randImage("FN.jpg", "WQ.jpg", "WQ1.jpg"))
		})
	engine.OnFullMatchGroup([]string{"早安", "早哇", "早上好", "ohayo", "哦哈哟", "お早う", "早好", "早", "早早早"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			process.SleepAbout1sTo2s()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"zzz...好涩哦..zzz....",
					"别...不要..zzz..那..zzz..",
					"嘻嘻..zzz..呐~..zzzz..",
					"...zzz....哧溜哧溜....",
				))
			case now >= 6 && now < 9:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"啊......早上好...(哈欠)",
					"唔......吧唧...早上...哈啊啊~~~\n早上好......",
					"早上好......",
					"早上好呜......呼啊啊~~~~",
					"啊......早上好。\n昨晚也很激情呢！",
					"吧唧吧唧......怎么了...已经早上了么...",
					"早上好！",
					"......看起来像是傍晚，其实已经早上了吗？",
					"早上好......欸~~~脸好近呢",
				))
			case now >= 9 && now < 18:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"哼！这个点还早啥，昨晚干啥去了！？",
					"熬夜了对吧熬夜了对吧熬夜了对吧？？？！",
					"是不是熬夜是不是熬夜是不是熬夜？！",
				))
			case now >= 18 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"早个啥？哼唧！我都准备洗洗睡了！",
					"不是...你看看几点了，哼！",
					"晚上好哇",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"中午好", "午安", "午好"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			if now > 11 && now < 15 { // 中午
				process.SleepAbout1sTo2s()
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安w",
					"午觉要好好睡哦，斯卡蒂会陪伴在你身旁的w",
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"睡你午觉去！哼唧！！",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"晚安", "oyasuminasai", "おやすみなさい", "晚好", "晚上好"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			process.SleepAbout1sTo2s()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"zzz...好涩哦..zzz....",
					"别...不要..zzz..那..zzz..",
					"嘻嘻..zzz..呐~..zzzz..",
					"...zzz....哧溜哧溜....",
				))
			case now >= 6 && now < 11:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"你可猝死算了吧！",
					"？啊这",
					"亲，这边建议赶快去睡觉呢~~~",
					"不可忍不可忍不可忍！！为何这还不猝死！！",
				))
			case now >= 11 && now < 15:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安w",
					"午觉要好好睡哦，斯卡蒂会陪伴在你身旁的w",
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"睡你午觉去！哼唧！！",
				))
			case now >= 15 && now < 19:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"难不成？？晚上不想睡觉？？现在休息",
					"就......挺离谱的...现在睡觉",
					"现在还是白天哦，睡觉还太早了",
				))
			case now >= 19 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"......(打瞌睡)",
					"呼...呼...已经睡着了哦~...呼......",
					"......我、我会在这守着你的，请务必好好睡着",
				))
			}
		})
	engine.OnKeywordGroup([]string{"任命助理", "观看作战记录","太棒了", "すごい", "sugoi", "斯国一", "よかった","头发","交谈"}, isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(randText(
				"为了你的安全着想，最好离我两米以上，也不要和我说话......我可不想你被冲着我来的危险伤到。",
				"看过我的履历了？我的过去可是很糟糕的。嗯，让我呆在身边，不仅你会被卷入麻烦，我也会很难办的。",
				"我的头发很长，很好看？啊，嗯，谢谢......要不要摸摸看？我的头发还是挺柔软的。这方面，我还算是有些自信的哦。",
				"我没看错吧？那个人，现在叫幽灵鲨是吗，怎么穿成了修女的样子？失忆了？这可麻烦了......博士，你得帮帮她，可不能让她忘记自己真正的职责。",
				"你这人，怎么这么执着，这样我不就只能老老实实保护你了吗。",
				"远方浩瀚蓝色与轻抚土地的白色碎片之下，埋葬着原本我所在乎的一切。我很害怕，害怕厄运会再次夺走我重要的人，所以我才会选择逃离，我不希望我珍视的人......因我而受伤。",
				"根据传说，我的族裔已经和那些灾祸战斗了无数年。说不定，我们也帮你们这些城市人把灾祸挡在了陆地之外......所以说，是不是该请我喝一杯，好好谢谢我？",
				"好，又度过了轻松的一天！没有会卷走队友的巨大触须，也没有蹲在角落里满手是血的疯狂敌人......光是上上战场什么的，对，已经很轻松了！",
				"睡着了？睡吧，博士，得做个干燥的好梦哟。",
				"斯卡蒂，赏金猎人。你真要签下我？我可是那种，会给你带来灾祸的人哦。",
				"嘿嘿，我的高性能发挥出来啦♪",
				"呯呯地冲过去，咚地打倒敌人，啪地解决掉——原来你们是这么战斗的，我记住啦。",
				"这样渴求我的力量，真的让我很困扰......如果你真的无所畏惧，那么做好准备，和我一起直面无边无际的黑暗吧。",
			))
		})
	engine.OnKeywordGroup([]string{"开始作战", "作战", "战斗开始", "行动出发", "行动开始", "部署", "行动结束"}, isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(randText(
				"我倒也挺擅长团队作战，只是......",
				"你是把大家的性命交到我手上了吗？",
				"对手是人类的话，终归还是没什么手感。",
				"和我面对过的灾厄相比，你们也太弱了。",
				"我习惯了独行。",
				"起舞吧。",
				"纠缠着我的噩梦啊，唱个歌吧。",
				"别眨眼，你会错过自己的死状。",
			))
		})

	engine.OnKeywordGroup([]string{"好吗", "是吗", "行不行", "能不能", "可不可以"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			if rand.Intn(2) == 0 {
				ctx.SendChain(randImage("YES.png", "NO.jpg"))
			}
		})
	engine.OnKeywordGroup([]string{"啊这"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			if rand.Intn(2) == 0 {
				ctx.SendChain(randImage("AZ.jpg", "AZ1.jpg"))
			}
		})
	engine.OnKeywordGroup([]string{"我好了"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText("啊......帽子！", "离我那么近......我，我可没有能完全保护你的自信！"))
		})
	engine.OnFullMatchGroup([]string{"？", "?", "¿"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)ん?", "ん？"))
			case 1, 2:
				ctx.SendChain(randImage("WH.jpg", "WH1.jpg", "WH2.jpg", "WH3.jpg"))
			}
		})
	engine.OnKeyword("离谱", isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)ん?", "ん？"))
			case 1, 2:
				ctx.SendChain(randImage("WH.jpg"))
			}
		})
	engine.OnKeyword("答应我", isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			process.SleepAbout1sTo2s()
			ctx.SendChain(randText("我无法回应你的请求"))
		})
}

func randText(text ...string) message.MessageSegment {
	return message.Text(text[rand.Intn(len(text))])
}

func randImage(file ...string) message.MessageSegment {
	return message.Image(res + file[rand.Intn(len(file))])
}

func randRecord(file ...string) message.MessageSegment {
	return message.Record(res + file[rand.Intn(len(file))])
}

// isAtriSleeping 凌晨0点到6点，斯卡蒂 在睡觉，不回应任何请求
func isAtriSleeping(ctx *zero.Ctx) bool {
	if now := time.Now().Hour(); now >= 1 && now < 6 {
		return false
	}
	return true
}
