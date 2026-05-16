package stdcommands

import (
	"github.com/xaelistic/yagpdb/v2/bot"
	"github.com/xaelistic/yagpdb/v2/bot/eventsystem"
	"github.com/xaelistic/yagpdb/v2/commands"
	"github.com/xaelistic/yagpdb/v2/common"
	"github.com/xaelistic/yagpdb/v2/stdcommands/advice"
	"github.com/xaelistic/yagpdb/v2/stdcommands/allocstat"
	"github.com/xaelistic/yagpdb/v2/stdcommands/banserver"
	"github.com/xaelistic/yagpdb/v2/stdcommands/calc"
	"github.com/xaelistic/yagpdb/v2/stdcommands/catfact"
	"github.com/xaelistic/yagpdb/v2/stdcommands/ccreqs"
	"github.com/xaelistic/yagpdb/v2/stdcommands/cleardm"
	"github.com/xaelistic/yagpdb/v2/stdcommands/createinvite"
	"github.com/xaelistic/yagpdb/v2/stdcommands/currentshard"
	"github.com/xaelistic/yagpdb/v2/stdcommands/currenttime"
	"github.com/xaelistic/yagpdb/v2/stdcommands/customembed"
	"github.com/xaelistic/yagpdb/v2/stdcommands/dadjoke"
	"github.com/xaelistic/yagpdb/v2/stdcommands/dcallvoice"
	"github.com/xaelistic/yagpdb/v2/stdcommands/define"
	"github.com/xaelistic/yagpdb/v2/stdcommands/dictionary"
	"github.com/xaelistic/yagpdb/v2/stdcommands/dogfact"
	"github.com/xaelistic/yagpdb/v2/stdcommands/eightball"
	"github.com/xaelistic/yagpdb/v2/stdcommands/findserver"
	"github.com/xaelistic/yagpdb/v2/stdcommands/forex"
	"github.com/xaelistic/yagpdb/v2/stdcommands/globalrl"
	"github.com/xaelistic/yagpdb/v2/stdcommands/guildunavailable"
	"github.com/xaelistic/yagpdb/v2/stdcommands/howlongtobeat"
	"github.com/xaelistic/yagpdb/v2/stdcommands/info"
	"github.com/xaelistic/yagpdb/v2/stdcommands/inspire"
	"github.com/xaelistic/yagpdb/v2/stdcommands/invite"
	"github.com/xaelistic/yagpdb/v2/stdcommands/leaveserver"
	"github.com/xaelistic/yagpdb/v2/stdcommands/listflags"
	"github.com/xaelistic/yagpdb/v2/stdcommands/listroles"
	"github.com/xaelistic/yagpdb/v2/stdcommands/memstats"
	"github.com/xaelistic/yagpdb/v2/stdcommands/ping"
	"github.com/xaelistic/yagpdb/v2/stdcommands/poll"
	"github.com/xaelistic/yagpdb/v2/stdcommands/roast"
	"github.com/xaelistic/yagpdb/v2/stdcommands/roll"
	"github.com/xaelistic/yagpdb/v2/stdcommands/setstatus"
	"github.com/xaelistic/yagpdb/v2/stdcommands/simpleembed"
	"github.com/xaelistic/yagpdb/v2/stdcommands/sleep"
	"github.com/xaelistic/yagpdb/v2/stdcommands/statedbg"
	"github.com/xaelistic/yagpdb/v2/stdcommands/stateinfo"
	"github.com/xaelistic/yagpdb/v2/stdcommands/throw"
	"github.com/xaelistic/yagpdb/v2/stdcommands/toggledbg"
	"github.com/xaelistic/yagpdb/v2/stdcommands/topcommands"
	"github.com/xaelistic/yagpdb/v2/stdcommands/topevents"
	"github.com/xaelistic/yagpdb/v2/stdcommands/topgames"
	"github.com/xaelistic/yagpdb/v2/stdcommands/topic"
	"github.com/xaelistic/yagpdb/v2/stdcommands/topservers"
	"github.com/xaelistic/yagpdb/v2/stdcommands/unbanserver"
	"github.com/xaelistic/yagpdb/v2/stdcommands/undelete"
	"github.com/xaelistic/yagpdb/v2/stdcommands/viewperms"
	"github.com/xaelistic/yagpdb/v2/stdcommands/weather"
	"github.com/xaelistic/yagpdb/v2/stdcommands/wouldyourather"
	"github.com/xaelistic/yagpdb/v2/stdcommands/xkcd"
	"github.com/xaelistic/yagpdb/v2/stdcommands/yagstatus"
)

var (
	_ bot.BotInitHandler       = (*Plugin)(nil)
	_ commands.CommandProvider = (*Plugin)(nil)
)

type Plugin struct{}

func (p *Plugin) PluginInfo() *common.PluginInfo {
	return &common.PluginInfo{
		Name:     "Standard Commands",
		SysName:  "standard_commands",
		Category: common.PluginCategoryCore,
	}
}

func (p *Plugin) AddCommands() {
	commands.AddRootCommands(p,
		// Info
		info.Command,
		invite.Command,

		// Standard
		define.Command,
		weather.Command,
		calc.Command,
		topic.Command,
		catfact.Command,
		dadjoke.Command,
		dogfact.Command,
		advice.Command,
		ping.Command,
		throw.Command,
		roll.Command,
		customembed.Command,
		simpleembed.Command,
		currenttime.Command,
		listroles.Command,
		memstats.Command,
		wouldyourather.Command,
		poll.Command,
		undelete.Command,
		viewperms.Command,
		topgames.Command,
		xkcd.Command,
		howlongtobeat.Command,
		inspire.Command,
		forex.Command,
		roast.Command,
		eightball.Command,

		// Maintenance
		stateinfo.Command,
		leaveserver.Command,
		banserver.Command,
		cleardm.Command,
		allocstat.Command,
		unbanserver.Command,
		topservers.Command,
		topcommands.Command,
		topevents.Command,
		currentshard.Command,
		guildunavailable.Command,
		yagstatus.Command,
		setstatus.Command,
		createinvite.Command,
		findserver.Command,
		dcallvoice.Command,
		ccreqs.Command,
		sleep.Command,
		toggledbg.Command,
		globalrl.Command,
		listflags.Command,
	)

	statedbg.Commands()
	commands.AddRootCommands(p, dictionary.Command)
}

func (p *Plugin) BotInit() {
	eventsystem.AddHandlerAsyncLastLegacy(p, ping.HandleMessageCreate, eventsystem.EventMessageCreate)
}

func RegisterPlugin() {
	common.RegisterPlugin(&Plugin{})
}
