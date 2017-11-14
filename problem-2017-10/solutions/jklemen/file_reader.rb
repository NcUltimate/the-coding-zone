module Intersections  
  class FileReader
    class << self
      def read!(filename)
        self.new(filename).read!
      end
    end

    attr_accessor :filename
    def initialize(filename)
      @filename = filename
    end

    def read!
      [
        Polygon.new(points1),
        Polygon.new(points2)
      ]
    end

    private

    def lines
      @lines ||= File.open(filename, 'r') do |f|
        f.readlines.map(&:chomp)
      end
    end

    def to_point(str)
      m = str.match(/(\d+) (\d+)/)
      x = m ? m[1].to_i : 0
      y = m ? m[2].to_i : 0
      Point.new(x, y)
    end

    def size1
      lines[0].to_i
    end

    def size2
      lines[size1 + 1].to_i
    end

    def points1
      @points1 ||= lines[1..size1].map(&method(:to_point))
    end

    def points2
      @points2 ||= lines[(size1 + 2)..-1].map(&method(:to_point))
    end
  end
end