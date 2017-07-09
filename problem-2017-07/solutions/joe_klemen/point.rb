module Mario
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

    def hash
      [x, y].hash
    end

    def eql?(p)
      self.==(p)
    end

    def equal?(p)
      self.==(p)
    end

    def inspect
      "(#{x}, #{y})"
    end

    def to_s
      inspect
    end
  end
end