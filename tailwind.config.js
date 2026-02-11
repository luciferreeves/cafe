export default {
    content: [
        "./templates/**/*.django",
        "./static/js/**/*.js",
    ],
    theme: {
        extend: {
            colors: {
                dark: 'var(--color-dark)',
                overlay: 'var(--color-overlay)',
                muted: 'var(--color-muted)',
                primary: {
                    DEFAULT: 'var(--color-primary)',
                    hover: 'var(--color-primary-hover)',
                },
                secondary: {
                    DEFAULT: 'var(--color-secondary)',
                    hover: 'var(--color-secondary-hover)',
                },
                warning: {
                    DEFAULT: 'var(--color-warning)',
                    hover: 'var(--color-warning-hover)',
                },
                success: {
                    DEFAULT: 'var(--color-success)',
                    hover: 'var(--color-success-hover)',
                },
                error: {
                    DEFAULT: 'var(--color-error)',
                    hover: 'var(--color-error-hover)',
                },
            },
            textColor: {
                DEFAULT: 'var(--color-text)',
                muted: 'var(--color-muted)',
            },
        },
    },
};
