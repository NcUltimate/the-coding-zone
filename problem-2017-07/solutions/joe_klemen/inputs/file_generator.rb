config = {
  width: 3 * 10000 - 2,
  height: 14,
  file: 'recursion_breaker.txt' 
}

File.open(config[:file], 'w') do |file|
  file.puts "#{config[:width]} #{config[:height]}"
  file.puts '0 0'

  # Place the blocks
  x, y = 1, 2
  until (config[:width] - x - 1).abs <= 2 && (config[:height] - y - 1).abs <= 2
    file.puts "#{x} #{y} B"

    if y == config[:height] - 3
      y = 2
    else
      x += 1
      y += 3
    end
  end
  file.puts "#{x} #{y} B"

  # Place the coins
  x, y = 2, 3
  until (config[:width] - x - 1).abs <= 1 && config[:height] - 2 == y
    file.puts "#{x} #{y} C"
    file.puts "#{x} #{y+1} C"

    if y == config[:height] - 2
      y = 3
    else
      x += 1
      y += 3
    end
  end
  file.puts "#{x} #{y} C"
  file.puts "#{x} #{y+1} C"
end