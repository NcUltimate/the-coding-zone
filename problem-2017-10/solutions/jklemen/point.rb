module Intersections
  class Point < Shape
    attr_accessor :x, :y
    def initialize(x = 0, y = 0)
      self.x = x.round(4)
      self.y = y.round(4)
    end

    def +(point)
      Point.new(self.x + point.x, self.y + point.y)
    end

    def -(point)
      Point.new(self.x - point.x, self.y - point.y)
    end

    def *(point)
      Point.new(self.x * point.x, self.y * point.y)
    end

    def eql?(point)
      (x - point.x).abs < 0.0001 &&
        (y - point.y).abs < 0.0001
    end

    def hash
      [x, y].hash
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

    def dist_from(point)
      ys = (point.y - y) ** 2
      xs = (point.x - x) ** 2
      Math.sqrt(xs + ys)
    end

    def to_s
      xs = (x.to_i == x ? x.to_i : x)
      ys = (y.to_i == y ? y.to_i : y)
      "(#{xs}, #{ys})"
    end

    def to_output_s
      xs = (x.to_i == x ? x.to_i : x)
      ys = (y.to_i == y ? y.to_i : y)
      "#{xs} #{ys}"
    end

    def inspect
      to_s
    end
  end
end