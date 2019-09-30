local M = {}

-- These can be set up in the main menu
M.max_score = 7
M.powerups = true
M.powerups_secret = false
M.curses = true
M.curses_secret = true
M.curses_affect = "ENEMY"
M.chain_type = "NONE"
M.chain_length = 15
M.tooltips = true
M.gravity = 20
M.jet_power_upper = 1250
M.jet_power_lower = 1000
M.arena_floor = "spikes"
M.arena_ceiling = "wall"
M.arena_wall_left = "wall"
M.arena_wall_right = "wall"

-- These are hardcoded at this point
M.pickup_time = 9
M.pickup_frequency_max = 5
M.pickup_frequency_min = 3
M.powerup_types = {hash("jet"), hash("zero_g"), hash("shield")}
M.curse_types = {hash("spin"), hash("confusion")}
M.respawn_delay = 2

M.pickup_boost_duration = 5
M.pickup_boost_multiplier = 1.5
M.pickup_zero_g_duration = 7
M.pickup_confusion_duration = 6
M.pickup_shield_duration = 5

M.curse_spin_force = 1000

return M