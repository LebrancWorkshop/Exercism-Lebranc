defmodule Lasagna do
  # Please define the 'expected_minutes_in_oven/0' function
  def expected_minutes_in_oven()
    40
  end

  # Please define the 'remaining_minutes_in_oven/1' function
  def remaining_minutes_in_oven(num)
    expected_minutes_in_oven() - num
  end

  # Please define the 'preparation_time_in_minutes/1' function
  def preparation_time_in_minutes(layer)
    layer * 2
  end

  # Please define the 'total_time_in_minutes/2' function
  def total_time_in_minutes(layer, time)
    time + (layer * 2)
  end

  # Please define the 'alarm/0' function
  def alarm()
    IO.puts("Ding!")
  end
end
