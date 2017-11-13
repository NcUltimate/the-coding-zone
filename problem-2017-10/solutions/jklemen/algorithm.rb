module Intersections
  class Algorithm
    def self.run!(grid, words)
      self.new(grid, words).run!
    end

    attr_accessor :grid, :words, :locations, :added
    def initialize(grid, words)
      self.grid = grid.map(&:upcase).map(&:chars)
      self.words = words.map(&:upcase)
      self.locations = {}
      self.added = 0
    end

    def run!
      puts "TODO"
    end
  end
end