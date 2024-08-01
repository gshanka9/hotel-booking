import React, { useEffect, useState } from 'react';
import axios from 'axios';
import {
  AppBar,
  Toolbar,
  Typography,
  Container,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Snackbar,
  CircularProgress,
  TextField,
  IconButton,
  Switch,
} from '@mui/material';
import { Search as SearchIcon, Brightness4 as DarkModeIcon, Brightness7 as LightModeIcon } from '@mui/icons-material';
import { createTheme, ThemeProvider } from '@mui/material/styles';

function App() {
  const [errors, setErrors] = useState([]);
  const [filteredErrors, setFilteredErrors] = useState([]);
  const [loading, setLoading] = useState(true);
  const [snackbarOpen, setSnackbarOpen] = useState(false);
  const [searchQuery, setSearchQuery] = useState('');
  const [darkMode, setDarkMode] = useState(false);

  useEffect(() => {
    const baseURL = 'http://localhost:8080/errors';

    axios.get(baseURL)
        .then(response => {
          if (response.status === 200 && response.data.length > 0) {
            setErrors(response.data);
            setFilteredErrors(response.data);
          } else {
            setSnackbarOpen(true);
          }
        })
        .catch(error => console.error('Error fetching log entries:', error))
        .finally(() => setLoading(false));
  }, []);

  const handleSearch = (event) => {
    const query = event.target.value.toLowerCase();
    setSearchQuery(query);
    const filtered = errors.filter(error =>
        error.msg.toLowerCase().includes(query) ||
        error.file.toLowerCase().includes(query) ||
        error.author.toLowerCase().includes(query)
    );
    setFilteredErrors(filtered);
  };

  const toggleDarkMode = () => {
    setDarkMode(!darkMode);
  };

  const theme = createTheme({
    palette: {
      mode: darkMode ? 'dark' : 'light',
    },
  });

  return (
      <ThemeProvider theme={theme}>
        <div className="App">
          <AppBar position="static">
            <Toolbar>
              <Typography variant="h6" style={{ flexGrow: 1 }}>
                Optum Error Log Dashboard
              </Typography>
              <Switch
                  checked={darkMode}
                  onChange={toggleDarkMode}
                  icon={<LightModeIcon />}
                  checkedIcon={<DarkModeIcon />}
              />
            </Toolbar>
          </AppBar>
          <Container style={{ marginTop: '20px' }}>
            <TextField
                variant="outlined"
                fullWidth
                placeholder="Search errors..."
                value={searchQuery}
                onChange={handleSearch}
                InputProps={{
                  startAdornment: <SearchIcon position="start" />,
                }}
            />
            {loading ? (
                <div style={{ textAlign: 'center', marginTop: '20px' }}>
                  <CircularProgress />
                  <Typography variant="body1">Loading...</Typography>
                </div>
            ) : (
                <TableContainer component={Paper} style={{ marginTop: '20px' }}>
                  <Table>
                    <TableHead>
                      <TableRow>
                        <TableCell>Error Message</TableCell>
                        <TableCell>File</TableCell>
                        <TableCell>Line Number</TableCell>
                        <TableCell>Author</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {filteredErrors.length > 0 ? (
                          filteredErrors.map((error, index) => (
                              <TableRow key={index}>
                                <TableCell>{error.msg}</TableCell>
                                <TableCell>{error.file}</TableCell>
                                <TableCell>{error.line}</TableCell>
                                <TableCell>{error.author}</TableCell>
                              </TableRow>
                          ))
                      ) : (
                          <TableRow>
                            <TableCell colSpan={4} style={{ textAlign: 'center' }}>
                              No errors found.
                            </TableCell>
                          </TableRow>
                      )}
                    </TableBody>
                  </Table>
                </TableContainer>
            )}
          </Container>
          <Snackbar
              open={snackbarOpen}
              autoHideDuration={6000}
              onClose={() => setSnackbarOpen(false)}
              message="No errors returned from the server."
          />
        </div>
      </ThemeProvider>
  );
}

export default App;
