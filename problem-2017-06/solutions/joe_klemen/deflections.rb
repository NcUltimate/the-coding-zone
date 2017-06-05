module LaserGrid
  class Point
    class << self
      def [](x, y)
        self.new(x, y)
      end
    end

    attr_accessor :x, :y
    def initialize(x, y)
      self.x = x
      self.y = y
    end

    def +(point)
      self.class[x + point.x, y + point.y]
    end

    def ==(point)
      x == point.x && y == point.y
    end

    def inspect
      "(#{x}, #{y})"
    end
  end

  class Grid
    attr_accessor :dimensions, :goal, :mirrors
    def initialize(dimensions, goal, mirrors = {})
      self.dimensions = dimensions
      self.goal = goal
      self.mirrors = mirrors
    end

    def type(point)
      return '*' if point == goal
      return '' if mirrors[point.x].nil?
      mirrors[point.x][point.y].to_s
    end

    def include?(point)
      point.x > -1 && point.x < width &&
        point.y > -1 && point.y < height
    end

    def width
      dimensions.x
    end

    def height
      dimensions.y
    end

    def size
      width * height
    end

    def inspect
      "#<Grid #{height}x#{width} goal: #{goal}>"
    end
  end

  class Algorithm
    class << self
      def run!(grid)
        self.new(grid).run!
      end
    end

    attr_accessor :grid
    def initialize(grid)
      self.grid = grid
    end

    def run!
      s1 = traverse(grid.goal, Point[-1, 0])
      s2 = traverse(grid.goal, Point[1, 0])
      s3 = traverse(grid.goal, Point[0, -1])
      s4 = traverse(grid.goal, Point[0, 1])
      [s1, s2, s3, s4].compact.map(&method(:pformat)).sort_by { |a| a }
    end

    def pformat(point)
      return ['N', point.y + 1] if point.x == -1
      return ['S', point.y + 1] if point.x == grid.width
      return ['E', point.x + 1] if point.y == grid.height
      return ['W', point.x + 1] if point.y == -1
    end

    private

    def traverse(prev, diff)
      visited = 1
      loop do
        return unless visited < grid.size

        new_point = prev + diff
        return if grid.type(new_point) == '*'
        return new_point unless grid.include?(new_point)

        diff =
          case grid.type(new_point)
          when '/'
            Point[-diff.y, -diff.x]
          when '\\'
            Point[diff.y, diff.x]
          else
            diff
          end
        visited += 1
        prev = new_point
      end
    end
  end

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
        LaserGrid::Point[*dimensions],
        LaserGrid::Point[*goal],
        mirrors
      ]
    end

    private

    def to_grid(coord)
      coord[0] = coord[0].to_i - 1
      coord[1] = coord[1].to_i - 1
      coord
    end

    def lines
      @lines ||= File.open(filename, 'r') do |f|
        f.readlines.map do |line|
          line.chomp.split(/\s+/)
        end
      end
    end

    def dimensions
      @dimensions ||= lines[0].map(&:to_i)
    end

    def goal
      @goal ||= to_grid(lines[1])
    end

    def mirrors
      @mirrors ||= lines[2..-1].each_with_object({}) do |loc, hash|
        x, y, mark = to_grid(loc)
        hash[x] ||= {}
        hash[x][y] = mark
      end
    end
  end
end

if ARGV.length >= 1
  inputs  = LaserGrid::FileReader.read!(ARGV[0])
  grid    = LaserGrid::Grid.new(*inputs)
  result  = LaserGrid::Algorithm.run!(grid)
  result.each { |r| puts r.join(' ') }
else
  puts 'Usage: ruby deflections.rb <filename>'
end
