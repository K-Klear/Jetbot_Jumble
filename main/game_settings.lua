local M = {}

-- These can be set up in the main menu
M.max_score = 3
M.powerups = true
M.powerups_secret = false
M.curses = true
M.curses_secret = false
M.curses_affect = "ENEMY"
M.chain_type = "NONE"
M.chain_length = 20
M.tooltips = false
M.gravity = 0
M.jet_power_upper = 1750
M.jet_power_lower = 1500
M.arena_floor = "spikes"
M.arena_ceiling = "bounce"
M.arena_wall_left = "wall"
M.arena_wall_right = "wall"

-- These are hardcoded at this point
M.pickup_time = 7
M.pickup_frequency_max = 5
M.pickup_frequency_min = 3
M.powerup_types = {hash("jet")}
M.curse_types = {hash("spin")}
M.respawn_delay = 2

M.pickup_boost_duration = 5
M.pickup_boost_multiplier = 1.5

M.curse_spin_force = 900

return M