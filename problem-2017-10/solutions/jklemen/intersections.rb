require 'colorize'
require_relative './file_reader.rb'
require_relative './point.rb'
require_relative './algorithm.rb'

if ARGV.length >= 1
  inputs  = WordSearch::FileReader.read!(ARGV[0])
  result  = WordSearch::Algorithm.run!(*inputs)
  print result
  puts
else
  puts 'Usage: ruby word_search.rb <filename>'
end
