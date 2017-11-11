# make sure the apt package lists are up to date
cookbook_file "apt-sources.list" do
  path "/etc/apt/sources.list"
end
execute 'apt_update' do
  command 'apt-get update'
end

# base configuration recipe in Chef.
package "wget"
package "ntp"
cookbook_file "ntp.conf" do
  path "/etc/ntp.conf"
end
execute 'ntp_restart' do
  command 'service ntp restart'
end

# Go installation
package "golang"
execute 'get-mux' do
  environment 'GOPATH' => '/go'
  command 'go get -u github.com/gorilla/mux'
end
execute 'get-pq' do
  environment 'GOPATH' => '/go'
  command 'go get -u github.com/lib/pq'
end

# postgres setup
package "postgresql"
execute 'postgres-setup' do
  command 'echo "CREATE DATABASE contactdb;" | sudo -u postgres psql'
end
execute 'postgres-password' do
  command 'echo "ALTER USER postgres WITH PASSWORD \'postgres\';" | sudo -u postgres psql'
end
execute 'database-setup' do
  cwd '/vagrant'
  command 'sudo -u postgres psql contactdb -f schema.sql'
end

# nginx setup
package "nginx"
execute 'nginx_pid' do
  command 'mkdir -p /run/nginx'
end
cookbook_file "nginx-config" do
  path "/etc/nginx/sites-available/default"
end
execute 'nginx_reload' do
  command 'nginx -s reload'
end

# install tmux and start the server in a background session
package "tmux"
execute 'create-server-session' do
  cwd '/vagrant'
  environment 'GOPATH' => '/go'
  command 'tmux new-session -d -s server'
end
execute 'start-server' do
  command "tmux send-keys -t server 'go run contactmanager.go' C-m"
end

