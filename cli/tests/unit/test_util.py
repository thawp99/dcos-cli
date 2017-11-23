import dcoscli.util as util


def test_choicelist():
    real_read = util._read_response

    def fake_read():
        return "3"

    util._read_response = fake_read
    result = util.choicelist(['der', 'die', 'das'])
    util._read_response = real_read

    assert result == 'das'


def test_choicelist_with_invalid_input():
    real_read = util._read_response

    def fake_read():
        return "invalid"

    util._read_response = fake_read
    result = util.choicelist(['der', 'die', 'das'])
    util._read_response = real_read

    assert result is None
