import React from 'react';
import { Button, Container, TextField, Typography } from '@mui/material';
import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

const App = () => {
  const [originalURL, setOriginalURL] = React.useState('');
  const [shortenedURL, setShortenedURL] = React.useState('');

  const handleShortenURL = async (e: any) => {
    e.preventDefault();
    console.log('Form submitted');
    try {
      const response = await axios.post(`${API_BASE_URL}/shorten`, {
        URL: originalURL,
      });
      setShortenedURL(response.data.short_url);
    } catch (error: any) {
      console.error('Error shortening URL:', error.message);
      // Handle error as needed
    }
  };

  return (
    <Container maxWidth="sm" style={{ marginTop: '2rem' }}>
      <Typography variant="h4" gutterBottom>
        URL Shortener
      </Typography>
      <form onSubmit={handleShortenURL}>
        <TextField
          label="Enter URL"
          fullWidth
          variant="outlined"
          margin="normal"
          value={originalURL}
          onChange={(e) => setOriginalURL(e.target.value)}
        />
        <Button type="submit" variant="contained" color="primary">
          Shorten URL
        </Button>
      </form>
      {shortenedURL && (
        <div style={{ marginTop: '1rem' }}>
          <Typography variant="h6" gutterBottom>
            Shortened URL:
          </Typography>
          <a href={shortenedURL} target="_blank" rel="noopener noreferrer">
            {shortenedURL}
          </a>
        </div>
      )}
    </Container>
  );
};

export default App;
