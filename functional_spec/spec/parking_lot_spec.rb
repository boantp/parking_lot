RSpec.describe 'Parking Lot', type: :aruba do
  let(:command) { run "parking_lot" }
  before(:each) { command.write("create_parking_lot 3\n") }

  it "can create a parking lot" do
    stop_all_commands
    expect(command.output).to end_with("Created a parking lot with 3 slots\n")
  end

  it "can park a car" do
    command.write("park KA-01-HH-3141 Black\n")
    stop_all_commands
    expect(command.output).to end_with("Allocated slot number: 1\n")
  end
  
  it "can unpark a car" do
    command.write("park KA-01-HH-3141 Black\n")
    command.write("leave 1\n")
    stop_all_commands
    expect(command.output).to end_with("Slot number 1 is free\n")
  end
  
  it "can report status" do
    command.write("park KA-01-HH-1234 White\n")
    command.write("park KA-01-HH-3141 Black\n")
    command.write("park KA-01-HH-9999 White\n")
    command.write("status\n")
    stop_all_commands
    expect(command.output).to end_with(<<-EOTXT
Slot No.    Registration No    Colour
1           KA-01-HH-1234      White
2           KA-01-HH-3141      Black
3           KA-01-HH-9999      White
EOTXT
)
  end
  
  pending "add more specs as needed"
end
