package provisioner

import "fmt"

func VhostConf(port,dns,projectName) string{
	return  fmt.Sprintf(`<VirtualHost *:%d>
   ServerAdmin webmaster@localhost
   ErrorLog ${APACHE_LOG_DIR}/error.log
   CustomLog ${APACHE_LOG_DIR}/access.log combined
   DocumentRoot /var/www/%s
   ServerName %s
   </VirtualHost>`,port,projectName,dns)
}

func ApacheConf(projectName string) string{
	return fmt.Sprintf(`<Directory /var/www/%s/>
        Options Indexes FollowSymLinks
        AllowOverride None
        Require all granted
        </Directory>
`,projectName)
}
