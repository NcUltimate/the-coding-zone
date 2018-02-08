target_value = ARGV[0].to_i
moves = 0
while target_value != 1
  target_value % 2 == 0 ? target_value /= 2 : target_value -= 1
  moves += 1
end

puts moves
