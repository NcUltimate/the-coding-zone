module Intersections
  class Polygon < Shape
    attr_accessor :points, :edges, :neighbors
    def initialize(points = [])
      self.neighbors = points.each_with_object({}) { |p,o| o[p]=Set.new }
      self.edges = compute_edges(points)
      self.points = Set.new(points)
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
      Set.new(intersections.compact)
    end

    def intersects_point?(point)
      return true if points.member?(point)
      return true if edges.any? { |e| e.intersects?(point) }
      
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

    def compute_edges(points)
      last = nil
      add_neighbor(points.last, points.first)
      initial = [Line.new(points.last, points.first)]
      points.reduce(initial) do |edges, p|
        if last.nil?
          last = p
          next edges
        end
        add_neighbor(last, p)
        edges << Line.new(last, p)
        last = p
        edges
      end
    end

    def add_neighbor(p1, p2)
      self.neighbors[p1] << p2
      self.neighbors[p2] << p1
    end
  end
end