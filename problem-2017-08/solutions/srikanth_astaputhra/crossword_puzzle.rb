require_relative "./wordsearch.rb"
if ARGV.length >= 1
  grid = Grid.new("input.txt")
  Wordsearch.search(grid)
else
  puts 'Please send file as argument'
end


