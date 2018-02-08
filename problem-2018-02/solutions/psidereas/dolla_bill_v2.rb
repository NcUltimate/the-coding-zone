#This is just a noisier version of V1

target_value = ARGV[0].to_i
x = target_value
moves = []
while x != 1
  if x % 2 == 0
    moves << "double"
    x /= 2
  else
    moves << "add"
    x -= 1
  end
end

puts "It'll take #{moves.count} moves to get to $#{target_value}"

puts "The values after each move are:"
start = 1
moves.count.times do
  action = moves.pop
  if action == "double"
    puts start * 2
    start *= 2
  else
    puts start + 1
    start += 1
  end
end
