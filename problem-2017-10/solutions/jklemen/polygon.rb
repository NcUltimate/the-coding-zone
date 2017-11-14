module Intersections
  class Polygon < Shape
    attr_accessor :points, :edges
    def initialize(points)
      self.points = points
      self.edges = compute_edges
    end

    def center
      @center ||= points.reduce(Point.new) do |center, point|
        center.x += point.x.to_f / points.length
        center.y += point.y.to_f / points.length
        center
      end
    end

    def intersections(line)
      intersections = edges.reduce([]) do |pts, edge|
        pts << edge.intersection(line)
      end
      intersections.compact.uniq
    end

    def intersects_point?(point)
      line = Line.new(point, point + Point.new(99999999, 0))
      intersections(line).length % 2 != 0
    end

    def intersects_line?(line)
      intersects_point?(p1) &&
        intersects_point?(p2) &&
        intersects_point?(midpoint)
    end

    def intersects_poly?(poly)
      points.any? { |p| poly.intersects_point?(p) } ||
        edges.any? { |e| poly.intersects_line?(e) }
    end

    private

    def compute_edges
      last = nil
      initial = [Line.new(points.last, points.first)]
      points.reduce(initial) do |edges, p|
        if last.nil?
          last = p
          next edges
        end
        edges << Line.new(last, p)
        last = p
        edges
      end
    end
  end
end