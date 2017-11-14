require_relative './file_reader.rb'
require_relative './shape.rb'
require_relative './point.rb'
require_relative './line.rb'
require_relative './polygon.rb'
require_relative './algorithm.rb'

if ARGV.length >= 1
  inputs = Intersections::FileReader.read!(ARGV[0])
  result = Intersections::Algorithm.run!(*inputs)
  print result
  puts
else
  puts 'Usage: ruby intersections.rb <filename>'
end
