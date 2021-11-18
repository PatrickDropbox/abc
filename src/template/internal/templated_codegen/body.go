package templated_codegen

import (
	fmt "fmt"
	io "io"
)

import (
	"github.com/pattyshack/abc/src/template/internal"
)

type Body struct {
	ind  string
	body []template.Statement
}

func (Body) Name() string { return "Body" }

func (template *Body) writeValue(
	output io.Writer, value interface{}, loc string) (int, error) {
	var valueBytes []byte
	switch val := value.(type) {
	case fmt.Stringer:
		valueBytes = []byte(val.String())
	case string:
		valueBytes = []byte(val)
	case []byte:
		valueBytes = val
	case bool:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case uint:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case uint8:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case uint16:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case uint32:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case uint64:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case int:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case int8:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case int16:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case int32:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case int64:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case float32:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case float64:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case complex64:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	case complex128:
		valueBytes = []byte(fmt.Sprintf("%v", val))
	default:
		return 0, fmt.Errorf("Unsupported output value type (%s): %v", loc, value)
	}

	return output.Write(valueBytes)
}

func (_template *Body) WriteTo(_output io.Writer) (int64, error) {
	_numWritten := int64(0)

	ind := _template.ind
	body := _template.body

	// body.template:14:0
	for _, statement := range body {

		// body.template:15:4
		switch stmt := statement.(type) {
		case *template.Atom:

			// body.template:18:8
			switch stmt.Id() {
			case template.CommentToken:

			case template.TextToken:

				// body.template:23:12
				{
					_n, _err := (&Text{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}

			case template.SubstitutionToken:

				// body.template:26:12
				{
					_n, _err := (&Substitute{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}

			case template.EmbedToken:

				// body.template:29:12
				{
					_n, _err := (&Embed{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}

			case template.CopySectionToken:

				// body.template:32:12
				{
					_n, _err := (&CopySection{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}

			case template.ContinueToken:

				// body.template:35:12
				{
					_n, _err := (&Continue{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}

			case template.BreakToken:

				// body.template:38:12
				{
					_n, _err := (&Break{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}

			case template.ReturnToken:

				// body.template:41:12
				{
					_n, _err := (&Return{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}

			case template.ErrorToken:

				// body.template:44:12
				{
					_n, _err := (&Error{ind, stmt}).WriteTo(_output)
					_numWritten += _n
					if _err != nil {
						return _numWritten, _err
					}
				}

			default:

				// body.template:46:20
				{
					_n, _err := _output.Write([]byte(`
            // `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}

				// body.template:47:15
				{
					_n, _err := _template.writeValue(
						_output,
						(statement.Loc()),
						"body.template:47:15")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}

				// body.template:47:33
				{
					_n, _err := _output.Write([]byte(`
            COMPILE ERROR: bug in template generation code
            Unexpected atom type: `))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}

				// body.template:49:34
				{
					_n, _err := _template.writeValue(
						_output,
						(stmt.Id()),
						"body.template:49:34")
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}

				// body.template:49:46
				{
					_n, _err := _output.Write([]byte(`
`))
					_numWritten += int64(_n)
					if _err != nil {
						return _numWritten, _err
					}
				}
			}

		case *template.For:

			// body.template:54:8
			{
				_n, _err := (&For{ind, stmt}).WriteTo(_output)
				_numWritten += _n
				if _err != nil {
					return _numWritten, _err
				}
			}

		case *template.Switch:

			// body.template:57:8
			{
				_n, _err := (&Switch{ind, stmt}).WriteTo(_output)
				_numWritten += _n
				if _err != nil {
					return _numWritten, _err
				}
			}

		case *template.If:

			// body.template:60:8
			{
				_n, _err := (&If{ind, stmt}).WriteTo(_output)
				_numWritten += _n
				if _err != nil {
					return _numWritten, _err
				}
			}

		default:

			// body.template:62:16
			{
				_n, _err := _output.Write([]byte(`
        // `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// body.template:63:11
			{
				_n, _err := _template.writeValue(
					_output,
					(statement.Loc()),
					"body.template:63:11")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// body.template:63:29
			{
				_n, _err := _output.Write([]byte(`
        COMPILE ERROR: bug in template generation code
        Unexpected statement type: `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// body.template:65:35
			{
				_n, _err := _template.writeValue(
					_output,
					(statement.Id()),
					"body.template:65:35")
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}

			// body.template:65:52
			{
				_n, _err := _output.Write([]byte(`
    `))
				_numWritten += int64(_n)
				if _err != nil {
					return _numWritten, _err
				}
			}
		}

		// body.template:66:11
		{
			_n, _err := _output.Write([]byte(`
`))
			_numWritten += int64(_n)
			if _err != nil {
				return _numWritten, _err
			}
		}
	}

	// body.template:67:7
	{
		_n, _err := _output.Write([]byte(`
`))
		_numWritten += int64(_n)
		if _err != nil {
			return _numWritten, _err
		}
	}

	return _numWritten, nil
}
