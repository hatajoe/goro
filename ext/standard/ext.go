package standard

import (
	"math"

	"github.com/MagicalTux/gophp/core"
)

// WARNING: This file is auto-generated. DO NOT EDIT

func init() {
	core.RegisterExt(&core.Ext{
		Name: "standard",
		Functions: map[string]*core.ExtFunction{
			"abs":                      &core.ExtFunction{Func: mathAbs, Args: []*core.ExtFunctionArg{}},
			"acos":                     &core.ExtFunction{Func: mathAcos, Args: []*core.ExtFunctionArg{}},
			"acosh":                    &core.ExtFunction{Func: mathACosh, Args: []*core.ExtFunctionArg{}},
			"asin":                     &core.ExtFunction{Func: mathAsin, Args: []*core.ExtFunctionArg{}},
			"asinh":                    &core.ExtFunction{Func: mathAsinh, Args: []*core.ExtFunctionArg{}},
			"atan":                     &core.ExtFunction{Func: mathAtan, Args: []*core.ExtFunctionArg{}},
			"atan2":                    &core.ExtFunction{Func: mathAtan2, Args: []*core.ExtFunctionArg{}},
			"atanh":                    &core.ExtFunction{Func: mathAtanh, Args: []*core.ExtFunctionArg{}},
			"constant":                 &core.ExtFunction{Func: constant, Args: []*core.ExtFunctionArg{}},
			"cos":                      &core.ExtFunction{Func: mathCos, Args: []*core.ExtFunctionArg{}},
			"cosh":                     &core.ExtFunction{Func: mathCosh, Args: []*core.ExtFunctionArg{}},
			"deg2rad":                  &core.ExtFunction{Func: mathDeg2rad, Args: []*core.ExtFunctionArg{}},
			"die":                      &core.ExtFunction{Func: die, Args: []*core.ExtFunctionArg{}},
			"dl":                       &core.ExtFunction{Func: stdFuncDl, Args: []*core.ExtFunctionArg{}},
			"echo":                     &core.ExtFunction{Func: stdFuncEcho, Args: []*core.ExtFunctionArg{}},
			"eval":                     &core.ExtFunction{Func: stdFuncEval, Args: []*core.ExtFunctionArg{}},
			"exit":                     &core.ExtFunction{Func: exit, Args: []*core.ExtFunctionArg{}},
			"exp":                      &core.ExtFunction{Func: mathExp, Args: []*core.ExtFunctionArg{}},
			"expm1":                    &core.ExtFunction{Func: mathExpm1, Args: []*core.ExtFunctionArg{}},
			"extension_loaded":         &core.ExtFunction{Func: stdFunc, Args: []*core.ExtFunctionArg{}},
			"fmod":                     &core.ExtFunction{Func: mathFmod, Args: []*core.ExtFunctionArg{}},
			"function_exists":          &core.ExtFunction{Func: stdFuncFuncExists, Args: []*core.ExtFunctionArg{}},
			"gc_collect_cycles":        &core.ExtFunction{Func: stdFuncGcCollectCycles, Args: []*core.ExtFunctionArg{}},
			"gc_disable":               &core.ExtFunction{Func: stdFuncGcDisable, Args: []*core.ExtFunctionArg{}},
			"gc_enable":                &core.ExtFunction{Func: stdFuncGcEnable, Args: []*core.ExtFunctionArg{}},
			"gc_enabled":               &core.ExtFunction{Func: stdFuncGcEnabled, Args: []*core.ExtFunctionArg{}},
			"gc_mem_caches":            &core.ExtFunction{Func: stdFuncGcMemCaches, Args: []*core.ExtFunctionArg{}},
			"get_cfg_var":              &core.ExtFunction{Func: stdFuncGetCfgVar, Args: []*core.ExtFunctionArg{}},
			"get_magic_quotes_gpc":     &core.ExtFunction{Func: getMagicQuotesGpc, Args: []*core.ExtFunctionArg{}},
			"get_magic_quotes_runtime": &core.ExtFunction{Func: getMagicQuotesRuntime, Args: []*core.ExtFunctionArg{}},
			"hrtime":                   &core.ExtFunction{Func: stdFuncHrTime, Args: []*core.ExtFunctionArg{}},
			"hypot":                    &core.ExtFunction{Func: mathHypot, Args: []*core.ExtFunctionArg{}},
			"php_sapi_name":            &core.ExtFunction{Func: stdFuncSapiName, Args: []*core.ExtFunctionArg{}},
			"php_uname":                &core.ExtFunction{Func: fncUname, Args: []*core.ExtFunctionArg{}},
			"phpversion":               &core.ExtFunction{Func: stdFuncPhpVersion, Args: []*core.ExtFunctionArg{}},
			"pi":                       &core.ExtFunction{Func: mathPi, Args: []*core.ExtFunctionArg{}},
			"sleep":                    &core.ExtFunction{Func: stdFuncSleep, Args: []*core.ExtFunctionArg{}},
			"str_replace":              &core.ExtFunction{Func: stdStrReplace, Args: []*core.ExtFunctionArg{}},
			"usleep":                   &core.ExtFunction{Func: stdFuncUsleep, Args: []*core.ExtFunctionArg{}},
			"var_dump":                 &core.ExtFunction{Func: stdFuncVarDump, Args: []*core.ExtFunctionArg{}},
			"zend_version":             &core.ExtFunction{Func: stdFuncZendVersion, Args: []*core.ExtFunctionArg{}},
		},
		Constants: map[core.ZString]*core.ZVal{
			"INF":                 core.ZFloat(math.Inf(0)).ZVal(),
			"M_1_PI":              core.ZFloat(1 / math.Pi).ZVal(),
			"M_2_PI":              core.ZFloat(2 / math.Pi).ZVal(),
			"M_2_SQRTPI":          core.ZFloat(2 / math.Sqrt(math.Pi)).ZVal(),
			"M_E":                 core.ZFloat(math.E).ZVal(),
			"M_EULER":             core.ZFloat(0.57721566490153286061).ZVal(),
			"M_LN2":               core.ZFloat(math.Ln2).ZVal(),
			"M_LNPI":              core.ZFloat(math.Log(math.Pi)).ZVal(),
			"M_LOG10E":            core.ZFloat(math.Log10E).ZVal(),
			"M_LOG2E":             core.ZFloat(math.Log2E).ZVal(),
			"M_PHI":               core.ZFloat(math.Phi).ZVal(),
			"M_PI":                core.ZFloat(math.Pi).ZVal(),
			"M_PI_2":              core.ZFloat(math.Pi / 2).ZVal(),
			"M_PI_4":              core.ZFloat(math.Pi / 4).ZVal(),
			"M_SQRT1_2":           core.ZFloat(1 / math.Sqrt(2)).ZVal(),
			"M_SQRT2":             core.ZFloat(math.Sqrt(2)).ZVal(),
			"M_SQRT3":             core.ZFloat(math.Sqrt(3)).ZVal(),
			"M_SQRTPI":            core.ZFloat(math.Sqrt(math.Pi)).ZVal(),
			"NAN":                 core.ZFloat(math.NaN()).ZVal(),
			"PHP_ROUND_HALF_DOWN": core.ZInt(2).ZVal(),
			"PHP_ROUND_HALF_EVEN": core.ZInt(3).ZVal(),
			"PHP_ROUND_HALF_ODD":  core.ZInt(4).ZVal(),
			"PHP_ROUND_HALF_UP":   core.ZInt(1).ZVal(),
			"PHP_VERSION":         core.ZString(core.VERSION).ZVal(),
		},
	})
}
