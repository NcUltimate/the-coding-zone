require_relative './move.rb'
require_relative './pathfinder.rb'
require_relative './point.rb'
require_relative './file_reader.rb'
require_relative './algorithm.rb'
require_relative './grid.rb'

if ARGV.length >= 1
  inputs  = Mario::FileReader.read!(ARGV[0])
  grid    = Mario::Grid.new(*inputs)
  result  = Mario::Algorithm.run!(grid)
  p result
else
  puts 'Usage: ruby mario.rb <filename>'
end
