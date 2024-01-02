require 'sinatra/base'
require 'mysql2'
require 'mysql2-cs-bind'
require 'erubis'

module Ishocon2
  class AuthenticationError < StandardError; end
  class PermissionDenied < StandardError; end
end

class Ishocon2::WebApp < Sinatra::Base
  session_secret = ENV['ISHOCON2_SESSION_SECRET'] || 'showwin_happy'
  use Rack::Session::Cookie, key: 'rack.session', secret: session_secret
  set :erb, escape_html: true
  set :public_folder, File.expand_path('../public', __FILE__)
  set :protection, true

  helpers do
    def config
      @config ||= {
        db: {
          host: ENV['ISHOCON2_DB_HOST'] || 'localhost',
          port: ENV['ISHOCON2_DB_PORT'] && ENV['ISHOCON2_DB_PORT'].to_i,
          username: ENV['ISHOCON2_DB_USER'] || 'ishocon',
          password: ENV['ISHOCON2_DB_PASSWORD'] || 'ishocon',
          database: ENV['ISHOCON2_DB_NAME'] || 'ishocon2'
        }
      }
    end

    def db
      return Thread.current[:ishocon2_db] if Thread.current[:ishocon2_db]
      client = Mysql2::Client.new(
        host: config[:db][:host],
        port: config[:db][:port],
        username: config[:db][:username],
        password: config[:db][:password],
        database: config[:db][:database],
        reconnect: true
      )
      client.query_options.merge!(symbolize_keys: true)
      Thread.current[:ishocon2_db] = client
      client
    end

    # def candidates
    #   @candidates ||= db.query('SELECT * FROM candidates').map do |row|
    #     {
    #       id: row[:id],
    #       name: row[:name],
    #       political_party: row[:political_party],
    #       sex: row[:sex],
    #     }
    #   end
    # end

    def election_results
      query = <<SQL
SELECT c.id, c.name, c.political_party, c.sex, v.count
FROM candidates AS c
LEFT OUTER JOIN
  (SELECT candidate_id, COUNT(*) AS count
  FROM votes
  GROUP BY candidate_id) AS v
ON c.id = v.candidate_id
ORDER BY v.count DESC
SQL
      db.xquery(query)
    end

    def voice_of_supporter(candidate_ids)
      query = <<SQL
SELECT keyword
FROM votes
WHERE candidate_id IN (?)
GROUP BY keyword
ORDER BY COUNT(*) DESC
LIMIT 10
SQL
      db.xquery(query, candidate_ids).map { |a| a[:keyword] }
    end

    def db_initialize
      db.query('DELETE FROM votes')
    end
  end

  get '/' do
    cs = []
    election_results.each_with_index do |r, i|
      # 上位10人と最下位のみ表示
      cs.push(r) if i < 10 || 28 < i
    end

    # parties_set = cs.map { |c| c[:political_party] }.uniq
    parties_set = db.query('SELECT political_party FROM candidates GROUP BY political_party')
    parties = {}
    parties_set.each { |a| parties[a[:political_party]] = 0 }
    election_results.each do |r|
      parties[r[:political_party]] += r[:count] || 0
    end

    sex_ratio = { '男': 0, '女': 0 }
    election_results.each do |r|
      sex_ratio[r[:sex].to_sym] += r[:count] || 0
    end

    erb :index, locals: { candidates: cs,
                          parties: parties,
                          sex_ratio: sex_ratio }
  end

  get '/candidates/:id' do
    candidate = db.xquery('SELECT * FROM candidates WHERE id = ?', params[:id]).first
    # candidate = candidates.first { |c| c[:id] == param[:id] }
    return redirect '/' if candidate.nil?
    votes = db.xquery('SELECT COUNT(*) AS count FROM votes WHERE candidate_id = ?', params[:id]).first[:count]
    keywords = voice_of_supporter([params[:id]])
    erb :candidate, locals: { candidate: candidate,
                              votes: votes,
                              keywords: keywords }
  end

  get '/political_parties/:name' do
    votes = 0
    election_results.each do |r|
      votes += r[:count] || 0 if r[:political_party] == params[:name]
    end
    candidates = db.xquery('SELECT * FROM candidates WHERE political_party = ?', params[:name])
    candidate_ids = candidates.map { |c| c[:id] }
    keywords = voice_of_supporter(candidate_ids)
    erb :political_party, locals: { political_party: params[:name],
                                    votes: votes,
                                    candidates: candidates,
                                    keywords: keywords }
  end

  get '/vote' do
    candidates = db.query('SELECT * FROM candidates')
    erb :vote, locals: { candidates: candidates, message: '' }
  end

  post '/vote' do
    user = db.xquery('SELECT * FROM users WHERE name = ? AND address = ? AND mynumber = ?',
                     params[:name],
                     params[:address],
                     params[:mynumber]).first
    candidate = db.xquery('SELECT * FROM candidates WHERE name = ?', params[:candidate]).first
    # candidate = candidates.first { |c| c[:name] == params[:candidate] }
    voted_count =
      user.nil? ? 0 : db.xquery('SELECT COUNT(*) AS count FROM votes WHERE user_id = ?', user[:id]).first[:count]

    candidates = db.query('SELECT * FROM candidates')
    cs = candidates
    if user.nil?
      return erb :vote, locals: { candidates: cs, message: '個人情報に誤りがあります' }
    elsif user[:votes] < (params[:vote_count].to_i + voted_count)
      return erb :vote, locals: { candidates: cs, message: '投票数が上限を超えています' }
    elsif params[:candidate].nil? || params[:candidate] == ''
      return erb :vote, locals: { candidates: cs, message: '候補者を記入してください' }
    elsif candidate.nil?
      return erb :vote, locals: { candidates: cs, message: '候補者を正しく記入してください' }
    elsif params[:keyword].nil? || params[:keyword] == ''
      return erb :vote, locals: { candidates: cs, message: '投票理由を記入してください' }
    end

    params[:vote_count].to_i.times do
      result = db.xquery('INSERT INTO votes (user_id, candidate_id, keyword) VALUES (?, ?, ?)',
                user[:id],
                candidate[:id],
                params[:keyword])
    end
    return erb :vote, locals: { candidates: cs, message: '投票に成功しました' }
  end

  get '/initialize' do
    db_initialize
  end

  get '/health' do
  end
end
