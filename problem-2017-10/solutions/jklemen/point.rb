module Intersections
  class Point < Shape
    attr_accessor :x, :y
    def initialize(x = 0, y = 0)
      self.x = x
      self.y = y
    end

    def +(point)
      Point.new(self.x + point.x, self.y + point.y)
    end

    def *(point)
      Point.new(self.x * point.x, self.y * point.y)
    end

    def eql?(point)
      (x - point.x).abs < 0.0001 &&
        (y - point.y).abs < 0.0001
    end

    def hash
      [x.round(4), y.round(4)].hash
    end

    def intersects_point?(point)
      eql?(point)
    end

    def intersects_line?(line)
      line.intersects_point?(self)
    end

    def intersects_poly?(poly)
      poly.intersects_point?(self)
    end

    def to_s
      "#<Point (#{x}, #{y})>"
    end

    def inspect
      to_s
    end
  end
end