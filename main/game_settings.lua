local M = {}

-- These can be set up in the main menu
M.max_score = 3
M.powerups = false
M.powerups_secret = false
M.curses = false
M.curses_secret = true
M.curses_affect = "ENEMY"
M.chain_type = "NONE"
M.chain_length = 20
M.tooltips = false
M.gravity = 10

-- These are hardcoded as this point
M.pickup_time
M.pickup_frequency_max = 8
M.pickup_frequency_min = 5


return M